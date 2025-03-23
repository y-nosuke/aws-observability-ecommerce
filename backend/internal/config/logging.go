package config

import "strings"

// LoggingPreset はログ設定のプリセットを定義します
type LoggingPreset struct {
	Level               string
	Format              string
	UseConsole          bool
	UseFile             bool
	LogFilePath         string
	UseCloudWatch       bool
	CreateLogGroup      bool
	CloudWatchLogGroup  string
	CloudWatchFlushSecs int
	CloudWatchBatchSize int
}

// 定義済みのログプリセット
var (
	// ConsoleOnlyPreset はコンソールのみへの出力を行うプリセット
	ConsoleOnlyPreset = LoggingPreset{
		Level:      "info",
		Format:     "json",
		UseConsole: true,
	}

	// FileOnlyPreset はファイルのみへの出力を行うプリセット
	FileOnlyPreset = LoggingPreset{
		Level:       "info",
		Format:      "json",
		UseFile:     true,
		LogFilePath: "logs/app.log",
	}

	// CloudWatchOnlyPreset はCloudWatch Logsのみへの出力を行うプリセット
	CloudWatchOnlyPreset = LoggingPreset{
		Level:               "info",
		Format:              "json",
		UseCloudWatch:       true,
		CreateLogGroup:      true,
		CloudWatchLogGroup:  "/aws-observability-ecommerce/backend",
		CloudWatchFlushSecs: 5,
		CloudWatchBatchSize: 100,
	}

	// LocalDebugPreset はローカル開発用プリセット
	LocalDebugPreset = LoggingPreset{
		Level:               "debug",
		Format:              "json",
		UseConsole:          true,
		UseFile:             true,
		LogFilePath:         "logs/app.log",
		UseCloudWatch:       true,
		CreateLogGroup:      true,
		CloudWatchLogGroup:  "/aws-observability-ecommerce/backend",
		CloudWatchFlushSecs: 5,
		CloudWatchBatchSize: 100,
	}

	// CloudStandardPreset はクラウド環境用のスタンダードプリセット
	CloudStandardPreset = LoggingPreset{
		Level:               "info",
		Format:              "json",
		UseConsole:          true,
		UseCloudWatch:       true,
		CreateLogGroup:      true,
		CloudWatchLogGroup:  "/aws-observability-ecommerce/backend",
		CloudWatchFlushSecs: 5,
		CloudWatchBatchSize: 100,
	}

	// MinimalLogsPreset は最小限のログ出力プリセット
	MinimalLogsPreset = LoggingPreset{
		Level:      "warn",
		Format:     "json",
		UseConsole: true,
	}
)

// ログプリセットのマップ
var logPresets = map[string]LoggingPreset{
	"console-only":    ConsoleOnlyPreset,
	"file-only":       FileOnlyPreset,
	"cloudwatch-only": CloudWatchOnlyPreset,
	"local-debug":     LocalDebugPreset,
	"cloud-standard":  CloudStandardPreset,
	"minimal":         MinimalLogsPreset,
}

// GetLoggingPreset は名前でプリセットを取得します
func GetLoggingPreset(name string) (LoggingPreset, bool) {
	preset, exists := logPresets[name]
	return preset, exists
}

// RegisterLoggingPreset はカスタムプリセットを登録します
func RegisterLoggingPreset(name string, preset LoggingPreset) {
	logPresets[name] = preset
}

// ListAvailableLoggingPresets は利用可能なプリセット名のスライスを返します
func ListAvailableLoggingPresets() []string {
	presets := make([]string, 0, len(logPresets))
	for name := range logPresets {
		presets = append(presets, name)
	}
	return presets
}

// ParsePresetNames はカンマ区切りのプリセット名を分割します
func ParsePresetNames(presetNames string) []string {
	if presetNames == "" {
		return nil
	}

	var result []string
	for name := range strings.SplitSeq(presetNames, ",") {
		name = strings.TrimSpace(name)
		if name != "" {
			result = append(result, name)
		}
	}

	return result
}
