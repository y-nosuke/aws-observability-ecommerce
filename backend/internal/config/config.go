package config

import (
	"errors"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config はアプリケーション全体の設定を格納する構造体
type Config struct {
	App struct {
		Name        string
		Version     string
		Environment string
	}
	Server struct {
		Port string
	}
	Log struct {
		Level               string
		Format              string
		UseConsole          bool   // コンソール出力の有効/無効
		UseFile             bool   // ファイル出力の有効/無効
		LogFilePath         string // ログファイルの出力パス
		UseCloudWatch       bool
		CreateLogGroup      bool // ロググループが存在しない場合に作成するかどうか
		CloudWatchLogGroup  string
		CloudWatchFlushSecs int // フラッシュ間隔（秒）
		CloudWatchBatchSize int // バッチサイズ
	}
	AWS struct {
		UseLocalStack bool
		Region        string
		Endpoint      string
		AccessKey     string
		SecretKey     string
		Token         string
	}
}

// アプリケーション設定インスタンス
var (
	config Config
	App    = &config.App
	Server = &config.Server
	Log    = &config.Log
	AWS    = &config.AWS
)

func init() {
	// 構成のロード
	if err := Load(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
}

// Load は環境変数と設定ファイルから設定をロードします
func Load() error {
	// 環境変数のデフォルト値の設定
	viper.SetDefault("app.name", "aws-observability-ecommerce")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("server.port", "8080")

	// ログ関連のデフォルト設定
	viper.SetDefault("log.level", "info")                                                  // デフォルトはinfo
	viper.SetDefault("log.format", "json")                                                 // デフォルトはJSON
	viper.SetDefault("log.useConsole", true)                                               // デフォルトはコンソール出力有効
	viper.SetDefault("log.useFile", false)                                                 // デフォルトはファイル出力無効
	viper.SetDefault("log.logFilePath", "logs/app.log")                                    // デフォルトのログファイルパス
	viper.SetDefault("log.useCloudWatch", false)                                           // デフォルトはCloudWatch無効
	viper.SetDefault("log.createLogGroup", true)                                           // デフォルトは自動作成を有効化
	viper.SetDefault("log.cloudWatchLogGroup", "/aws-observability-ecommerce/dev/backend") // デフォルトのロググループ名
	viper.SetDefault("log.cloudWatchFlushSecs", 5)                                         // デフォルトは5秒間隔
	viper.SetDefault("log.cloudWatchBatchSize", 100)                                       // デフォルトは100件

	// AWS関連のデフォルト設定
	viper.SetDefault("aws.useLocalStack", false)               // デフォルトはLocalStack無効
	viper.SetDefault("aws.region", "ap-northeast-1")           // デフォルトリージョン
	viper.SetDefault("aws.endpoint", "http://localstack:4566") // LocalStack使用時のエンドポイント
	viper.SetDefault("aws.accessKey", "")                      // デフォルトは空文字列
	viper.SetDefault("aws.secretKey", "")                      // デフォルトは空文字列
	viper.SetDefault("aws.token", "")                          // デフォルトは空文字列

	// ログプリセットの適用（環境変数より先に処理）
	if err := viper.BindEnv("log.preset", "LOG_PRESET"); err != nil {
		return err
	}
	presetNames := viper.GetString("log.preset")
	if presetNames != "" {
		// 複数プリセットの適用（カンマ区切りで指定可能）
		applyLoggingPresets(ParsePresetNames(presetNames))
	}

	// 環境変数の読み込み
	viper.AutomaticEnv()
	// アプリケーション関連の環境変数にプレフィックスを設定
	viper.SetEnvPrefix("APP")

	// 環境変数のバインド
	if err := viper.BindEnv("app.name", "APP_NAME"); err != nil {
		return err
	}
	if err := viper.BindEnv("app.version", "APP_VERSION"); err != nil {
		return err
	}
	if err := viper.BindEnv("app.environment", "APP_ENV"); err != nil {
		return err
	}
	if err := viper.BindEnv("server.port", "PORT"); err != nil {
		return err
	}

	// ログ関連の環境変数のバインド
	if err := viper.BindEnv("log.level", "LOG_LEVEL"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.format", "LOG_FORMAT"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.useConsole", "LOG_USE_CONSOLE"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.useFile", "LOG_USE_FILE"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.logFilePath", "LOG_FILE_PATH"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.useCloudWatch", "USE_CLOUDWATCH"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.createLogGroup", "LOG_CREATE_LOG_GROUP"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.cloudWatchLogGroup", "LOG_GROUP_NAME"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.cloudWatchFlushSecs", "LOG_CLOUDWATCH_FLUSH_SECS"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.cloudWatchBatchSize", "LOG_CLOUDWATCH_BATCH_SIZE"); err != nil {
		return err
	}

	// AWS関連の環境変数のバインド
	if err := viper.BindEnv("aws.useLocalStack", "AWS_USE_LOCALSTACK"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.region", "AWS_REGION"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.endpoint", "AWS_ENDPOINT"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.accessKey", "AWS_ACCESS_KEY_ID"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.secretKey", "AWS_SECRET_ACCESS_KEY"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.token", "AWS_SESSION_TOKEN"); err != nil {
		return err
	}

	// 設定ファイルの読み込み（存在する場合）
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 設定ファイルの読み込み（存在しなくてもエラーとしない）
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	// viper.Unmarshalを使って設定を一括で読み込む
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}

// applyLoggingPresets は複数のプリセットを順番に適用します（後勝ち）
func applyLoggingPresets(presetNames []string) {
	// 適用したプリセットを記録
	var appliedPresets []string

	// 各プリセットを順番に適用（後勝ち）
	for _, presetName := range presetNames {
		preset, exists := GetLoggingPreset(presetName)
		if exists {
			// プリセットを適用
			applyPreset(preset)
			appliedPresets = append(appliedPresets, presetName)
		} else {
			// 警告ログを出力
			log.Printf("WARNING: Unknown logging preset '%s'. Skipping. Available presets: %v",
				presetName, ListAvailableLoggingPresets())
		}
	}

	if len(appliedPresets) > 0 {
		log.Printf("Applied logging presets: %v", strings.Join(appliedPresets, ", "))
	}
}

// applyPreset は単一のプリセットをviperに適用します
func applyPreset(preset LoggingPreset) {
	// プリセットの各フィールドを適用（空の値はスキップ）
	if preset.Level != "" {
		viper.Set("log.level", preset.Level)
	}

	if preset.Format != "" {
		viper.Set("log.format", preset.Format)
	}

	// ブール値とその他の値は常に設定（デフォルト値もあるため）
	viper.Set("log.useConsole", preset.UseConsole)
	viper.Set("log.useFile", preset.UseFile)

	if preset.LogFilePath != "" {
		viper.Set("log.logFilePath", preset.LogFilePath)
	}

	viper.Set("log.useCloudWatch", preset.UseCloudWatch)
	viper.Set("log.createLogGroup", preset.CreateLogGroup)

	if preset.CloudWatchLogGroup != "" {
		viper.Set("log.cloudWatchLogGroup", preset.CloudWatchLogGroup)
	}

	if preset.CloudWatchFlushSecs > 0 {
		viper.Set("log.cloudWatchFlushSecs", preset.CloudWatchFlushSecs)
	}

	if preset.CloudWatchBatchSize > 0 {
		viper.Set("log.cloudWatchBatchSize", preset.CloudWatchBatchSize)
	}
}
