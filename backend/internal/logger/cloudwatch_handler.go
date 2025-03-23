package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/google/uuid"
)

// CloudWatchHandler はCloudWatch Logsへのログ転送を行うslogハンドラー
type CloudWatchHandler struct {
	id             string // インスタンス識別用
	client         *cloudwatchlogs.Client
	logGroupName   string
	logStreamName  string
	sequenceToken  *string
	attrs          []slog.Attr
	level          slog.Level
	buffer         []types.InputLogEvent
	mu             sync.Mutex // バッファ用のミューテックス
	quit           chan struct{}
	wg             *sync.WaitGroup // ゴルーチン待機用
	createLogGroup bool            // ロググループ自動作成フラグ
	flushInterval  time.Duration
	batchSize      int
}

// CloudWatchHandlerOption はCloudWatchHandlerのオプション設定用関数型
type CloudWatchHandlerOption func(*CloudWatchHandler)

// WithLevel はハンドラーのログレベルを設定するオプション
func WithLevel(level slog.Level) CloudWatchHandlerOption {
	return func(h *CloudWatchHandler) {
		h.level = level
	}
}

// WithCreateLogGroup はロググループの自動作成を設定するオプション
func WithCreateLogGroup(create bool) CloudWatchHandlerOption {
	return func(h *CloudWatchHandler) {
		h.createLogGroup = create
	}
}

// WithFlushInterval はログのフラッシュ間隔を設定するオプション
func WithFlushInterval(interval time.Duration) CloudWatchHandlerOption {
	return func(h *CloudWatchHandler) {
		h.flushInterval = interval
	}
}

// WithBatchSize はログのバッチサイズを設定するオプション
func WithBatchSize(size int) CloudWatchHandlerOption {
	return func(h *CloudWatchHandler) {
		h.batchSize = size
	}
}

// NewCloudWatchHandler は新しいCloudWatchHandlerを作成します
func NewCloudWatchHandler(client *cloudwatchlogs.Client, logGroupName string, opts ...CloudWatchHandlerOption) (*CloudWatchHandler, error) {
	// インスタンス識別用のIDを生成
	id := uuid.New().String()

	// 一意のログストリーム名を生成
	logStreamName := fmt.Sprintf("app-%s", uuid.New().String())

	// 共有WaitGroupを作成
	wg := &sync.WaitGroup{}

	// デフォルト設定でハンドラーを初期化
	h := &CloudWatchHandler{
		id:             id, // 一意のIDを設定
		client:         client,
		logGroupName:   logGroupName,
		logStreamName:  logStreamName,
		level:          slog.LevelInfo, // デフォルトはINFOレベル
		buffer:         make([]types.InputLogEvent, 0, 100),
		quit:           make(chan struct{}),
		wg:             wg,
		createLogGroup: true,            // デフォルトは自動作成を有効化
		flushInterval:  5 * time.Second, // デフォルトは5秒間隔
		batchSize:      100,             // デフォルトは最大100件
	}

	// オプションを適用
	for _, opt := range opts {
		opt(h)
	}

	// デバッグ情報を出力
	fmt.Printf("[%s] Initializing CloudWatch Logs handler with log group: %s, stream: %s\n",
		h.id, logGroupName, logStreamName)

	// ロググループの存在確認
	ctx := context.Background()
	groups, err := client.DescribeLogGroups(ctx, &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: aws.String(logGroupName),
	})
	if err != nil {
		fmt.Printf("[%s] Error checking log group existence: %v\n", h.id, err)
		return nil, fmt.Errorf("error checking log group existence: %w", err)
	}

	// 既存のロググループをチェック
	logGroupExists := false
	for _, group := range groups.LogGroups {
		if *group.LogGroupName == logGroupName {
			logGroupExists = true
			fmt.Printf("[%s] Found existing log group: %s\n", h.id, logGroupName)
			break
		}
	}

	// 同名のロググループが存在しない場合
	if !logGroupExists {
		if h.createLogGroup {
			fmt.Printf("[%s] Log group not exists. Creating log group: %s\n", h.id, logGroupName)
			_, err = client.CreateLogGroup(ctx, &cloudwatchlogs.CreateLogGroupInput{
				LogGroupName: aws.String(logGroupName),
			})
			if err != nil {
				fmt.Printf("[%s] Failed to create log group: %v\n", h.id, err)
				return nil, fmt.Errorf("failed to create log group: %w", err)
			}
			fmt.Printf("[%s] Created log group: %s\n", h.id, logGroupName)
		} else {
			errMsg := fmt.Sprintf("log group %s does not exist and auto-creation is disabled", logGroupName)
			fmt.Printf("[%s] %s\n", h.id, errMsg)
			return nil, fmt.Errorf(errMsg)
		}
	}

	// ログストリームの作成
	_, err = client.CreateLogStream(ctx, &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String(logGroupName),
		LogStreamName: aws.String(logStreamName),
	})
	if err != nil {
		fmt.Printf("[%s] Failed to create log stream: %v\n", h.id, err)
		return nil, fmt.Errorf("failed to create log stream: %w", err)
	}
	fmt.Printf("[%s] Created log stream successfully\n", h.id)

	// 定期的なログフラッシュを行うゴルーチンを開始
	h.wg.Add(1)
	go h.flushPeriodically()

	return h, nil
}

// logStderr はメッセージを標準エラーに書き込み、エラーを返します
func logStderr(format string, args ...any) error {
	_, err := fmt.Fprintf(os.Stderr, format, args...)
	return err
}

// flushPeriodically は定期的にバッファされたログをCloudWatch Logsに送信します
func (h *CloudWatchHandler) flushPeriodically() {
	fmt.Printf("[%s] Starting flush goroutine\n", h.id)
	defer h.wg.Done()

	ticker := time.NewTicker(h.flushInterval)
	defer ticker.Stop()

	retryCount := 0
	maxRetries := 3

	for {
		select {
		case <-ticker.C:
			h.mu.Lock() // バッファチェック前にロックを取得
			bufLen := len(h.buffer)
			h.mu.Unlock() // すぐにアンロック

			fmt.Printf("[%s] Periodic flush check, current buffer size: %d\n", h.id, bufLen)

			if bufLen == 0 {
				continue // バッファが空の場合はスキップ
			}

			err := h.Flush()
			if err != nil {
				// エラーは標準エラー出力に記録
				if logErr := logStderr("[%s] Error flushing logs to CloudWatch: %v\n", h.id, err); logErr != nil {
					// 標準エラー出力への書き込みに失敗した場合は何もできないので続行
					fmt.Printf("[%s] Error flushing logs to CloudWatch: %v\n", h.id, err)
				}

				// リトライ処理
				if retryCount < maxRetries {
					retryCount++
					// 指数バックオフでリトライ
					retryTime := time.Duration(1<<retryCount) * time.Second
					if logErr := logStderr("[%s] Will retry in %v (attempt %d/%d)\n", h.id, retryTime, retryCount, maxRetries); logErr != nil {
						// 同様に、標準エラー出力への書き込みに失敗した場合は静かに無視
						fmt.Printf("[%s] Will retry in %v (attempt %d/%d)\n", h.id, retryTime, retryCount, maxRetries)
					}
					time.Sleep(retryTime)
					continue
				}

				// リトライ失敗後はローカルファイルにフォールバック
				if fallbackErr := h.fallbackToFile(err); fallbackErr != nil {
					if logErr := logStderr("[%s] Failed to save logs to fallback file: %v\n", h.id, fallbackErr); logErr != nil {
						// 同様に、標準エラー出力への書き込みに失敗した場合は静かに無視
						fmt.Printf("[%s] Failed to save logs to fallback file: %v\n", h.id, fallbackErr)
					}
				}
				retryCount = 0
				fmt.Printf("[%s] Flushed buffered logs after retry. %d\n", h.id, retryCount)
			} else {
				fmt.Printf("[%s] Flushed buffered logs successfully.\n", h.id)
				retryCount = 0
			}
		case <-h.quit:
			fmt.Printf("[%s] Received quit signal, flushing remaining logs\n", h.id)
			// 終了時に残りのログをフラッシュ
			h.mu.Lock()
			bufLen := len(h.buffer)
			h.mu.Unlock()

			if bufLen > 0 {
				if err := h.Flush(); err != nil {
					if logErr := logStderr("[%s] Error flushing logs on shutdown: %v\n", h.id, err); logErr != nil {
						fmt.Printf("[%s] Error flushing logs on shutdown: %v\n", h.id, err)
					}
					// 最後のフォールバック試行
					if fallbackErr := h.fallbackToFile(err); fallbackErr != nil {
						if logErr := logStderr("[%s] Failed to save final logs to fallback file: %v\n", h.id, fallbackErr); logErr != nil {
							fmt.Printf("[%s] Failed to save final logs to fallback file: %v\n", h.id, fallbackErr)
						}
					}
				}
			}
			fmt.Printf("[%s] Flush goroutine terminated\n", h.id)
			return
		}
	}
}

// fallbackToFile はCloudWatch Logsへの送信に失敗した場合にログをファイルに保存します
func (h *CloudWatchHandler) fallbackToFile(originalErr error) (err error) {
	fmt.Printf("[%s] Falling back to file\n", h.id)
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.buffer) == 0 {
		return nil
	}

	// フォールバック用のファイル名
	fallbackFile := fmt.Sprintf("logs/cloudwatch_fallback_%s_%s.log", h.id, time.Now().Format("20060102_150405"))

	// ディレクトリが存在することを確認
	if err = os.MkdirAll("logs", 0755); err != nil {
		return fmt.Errorf("failed to create fallback log directory: %w", err)
	}

	// ファイルを開く
	file, err := os.OpenFile(fallbackFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open fallback log file: %w", err)
	}
	defer func(file *os.File) {
		if closeErr := file.Close(); closeErr != nil {
			err = fmt.Errorf("original error: %v, failed to close fallback file: %w", err, closeErr)
		}
	}(file)

	// ヘッダー情報を書き込み
	if _, err := file.WriteString(fmt.Sprintf("# CloudWatch Logs fallback - %s\n", time.Now().Format(time.RFC3339))); err != nil {
		return fmt.Errorf("failed to write to fallback file: %w", err)
	}
	if _, err := file.WriteString(fmt.Sprintf("# Original error: %v\n", originalErr)); err != nil {
		return fmt.Errorf("failed to write to fallback file: %w", err)
	}
	if _, err := file.WriteString(fmt.Sprintf("# Log group: %s, Log stream: %s\n", h.logGroupName, h.logStreamName)); err != nil {
		return fmt.Errorf("failed to write to fallback file: %w", err)
	}
	if _, err := file.WriteString("---\n"); err != nil {
		return fmt.Errorf("failed to write to fallback file: %w", err)
	}

	// バッファ内のログイベントをファイルに書き込み
	for _, event := range h.buffer {
		timestamp := time.UnixMilli(*event.Timestamp).Format(time.RFC3339)
		if _, err := file.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, *event.Message)); err != nil {
			return fmt.Errorf("failed to write log event to fallback file: %w", err)
		}
	}

	// 標準出力にフォールバックファイルの情報を書き込み
	fmt.Printf("[%s] Logs saved to fallback file: %s\n", h.id, fallbackFile)

	// バッファをクリア
	h.buffer = h.buffer[:0]

	return nil
}

// Close はハンドラーを閉じ、残りのログをフラッシュします
func (h *CloudWatchHandler) Close() error {
	fmt.Printf("[%s] Closing handler\n", h.id)

	// 終了シグナルを送信
	close(h.quit)

	// ゴルーチンの終了を待機
	h.wg.Wait()

	fmt.Printf("[%s] Handler closed\n", h.id)
	return nil
}

// Flush はバッファされたログをCloudWatch Logsに送信します
func (h *CloudWatchHandler) Flush() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	bufLen := len(h.buffer)

	if bufLen == 0 {
		return nil
	}

	// ローカルコピーを作成して送信
	events := make([]types.InputLogEvent, bufLen)
	copy(events, h.buffer)

	// CloudWatch Logsへの送信準備
	input := &cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(h.logGroupName),
		LogStreamName: aws.String(h.logStreamName),
		LogEvents:     events,
	}

	// シーケンストークンがある場合は設定
	if h.sequenceToken != nil {
		input.SequenceToken = h.sequenceToken
	}

	// ログを送信
	resp, err := h.client.PutLogEvents(context.Background(), input)
	if err != nil {
		fmt.Printf("[%s] Failed to put log events: %v\n", h.id, err)
		return fmt.Errorf("failed to put log events: %w", err)
	}

	// 次回のリクエスト用にシーケンストークンを更新
	h.sequenceToken = resp.NextSequenceToken

	fmt.Printf("[%s] Successfully sent %d events to CloudWatch Logs\n", h.id, bufLen)

	// バッファをクリア
	h.buffer = h.buffer[:0]

	return nil
}

// Enabled はログレベルが設定されたレベル以上かどうかを返します
func (h *CloudWatchHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

// Handle はログレコードを処理します
func (h *CloudWatchHandler) Handle(ctx context.Context, record slog.Record) error {
	// ログレベルが設定レベル未満の場合は何もしない
	if !h.Enabled(ctx, record.Level) {
		return nil
	}

	// ログレコードを構造体に変換
	logMap := make(map[string]any)
	logMap["time"] = record.Time.Format(time.RFC3339)
	logMap["level"] = record.Level.String()
	logMap["message"] = record.Message
	logMap["handler_id"] = h.id // ハンドラーIDを追加

	// 属性を追加
	record.Attrs(func(attr slog.Attr) bool {
		addAttr(logMap, "", attr)
		return true
	})

	// グローバル属性を追加
	for _, attr := range h.attrs {
		addAttr(logMap, "", attr)
	}

	// ログをJSON形式にシリアライズ
	jsonData, err := json.Marshal(logMap)
	if err != nil {
		return err
	}

	// ログイベントを作成
	logEvent := types.InputLogEvent{
		Message:   aws.String(string(jsonData)),
		Timestamp: aws.Int64(record.Time.UnixMilli()),
	}

	// バッファにログを追加する前にロックを取得
	h.mu.Lock()
	h.buffer = append(h.buffer, logEvent)
	bufferLen := len(h.buffer)
	h.mu.Unlock()

	fmt.Printf("[%s] Added log event to buffer, current size: %d\n", h.id, bufferLen)

	// バッファサイズが閾値を超えたらフラッシュ
	if bufferLen >= h.batchSize {
		fmt.Printf("[%s] Buffer size reached batch threshold, flushing\n", h.id)
		return h.Flush()
	}

	return nil
}

// WithAttrs は属性を持つ新しいハンドラーを返します
func (h *CloudWatchHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	fmt.Printf("[%s] Creating new handler with additional attributes\n", h.id)

	// 属性を結合
	newAttrs := make([]slog.Attr, len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)
	copy(newAttrs[len(h.attrs):], attrs)

	// 新しいインスタンスID
	newId := uuid.New().String()

	// 新しいハンドラーを作成（バッファと終了用チャネルは共有しない）
	newHandler := &CloudWatchHandler{
		client:        h.client,
		logGroupName:  h.logGroupName,
		logStreamName: h.logStreamName,
		sequenceToken: h.sequenceToken,
		attrs:         newAttrs,
		level:         h.level,
		buffer:        make([]types.InputLogEvent, 0, 100), // 新しいバッファを作成
		flushInterval: h.flushInterval,
		batchSize:     h.batchSize,
		quit:          make(chan struct{}), // 新しい終了チャネル
		wg:            &sync.WaitGroup{},   // 新しいWaitGroup
		id:            newId,               // 新しいID
	}

	// 定期的なログフラッシュを行うゴルーチンを開始
	newHandler.wg.Add(1)
	go newHandler.flushPeriodically()

	fmt.Printf("[%s] Created new handler with ID: %s\n", h.id, newId)
	return newHandler
}

// WithGroup はグループ化された属性を持つ新しいハンドラーを返します
func (h *CloudWatchHandler) WithGroup(name string) slog.Handler {
	fmt.Printf("[%s] Creating new handler with group: %s\n", h.id, name)

	if name == "" {
		// 新しいハンドラーを作成して返す（コピーを避ける）
		return h.WithAttrs(nil)
	}

	// 属性をグループ化
	newAttrs := make([]slog.Attr, len(h.attrs))
	for i, attr := range h.attrs {
		if attr.Key != "" {
			attr.Key = name + "." + attr.Key
		}
		newAttrs[i] = attr
	}

	// 新しいインスタンスID
	newId := uuid.New().String()

	// 新しいハンドラーを作成
	newHandler := &CloudWatchHandler{
		client:        h.client,
		logGroupName:  h.logGroupName,
		logStreamName: h.logStreamName,
		sequenceToken: h.sequenceToken,
		attrs:         newAttrs,
		level:         h.level,
		buffer:        make([]types.InputLogEvent, 0, 100), // 新しいバッファを作成
		flushInterval: h.flushInterval,
		batchSize:     h.batchSize,
		quit:          make(chan struct{}), // 新しい終了チャネル
		wg:            &sync.WaitGroup{},   // 新しいWaitGroup
		id:            newId,               // 新しいID
	}

	// 定期的なログフラッシュを行うゴルーチンを開始
	newHandler.wg.Add(1)
	go newHandler.flushPeriodically()

	fmt.Printf("[%s] Created new handler with ID: %s and group: %s\n", h.id, newId, name)
	return newHandler
}

// addAttr はログマップに属性を追加するヘルパー関数
func addAttr(m map[string]any, prefix string, attr slog.Attr) {
	key := attr.Key
	if prefix != "" {
		key = prefix + "." + key
	}

	switch attr.Value.Kind() {
	case slog.KindBool, slog.KindInt64, slog.KindUint64, slog.KindFloat64, slog.KindString:
		m[key] = attr.Value.Any()
	case slog.KindTime:
		m[key] = attr.Value.Time().Format(time.RFC3339)
	case slog.KindDuration:
		m[key] = attr.Value.Duration().String()
	case slog.KindGroup:
		for _, a := range attr.Value.Group() {
			addAttr(m, key, a)
		}
	case slog.KindLogValuer:
		addAttr(m, prefix, slog.Attr{
			Key:   attr.Key,
			Value: attr.Value.LogValuer().LogValue(),
		})
	default:
		m[key] = fmt.Sprintf("%v", attr.Value)
	}
}
