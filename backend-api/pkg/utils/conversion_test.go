package utils

import (
	"math"
	"testing"
)

func TestSafeIntToInt32(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		expected  int32
		expectErr bool
	}{
		{
			name:      "正常値（正の数）",
			input:     100,
			expected:  100,
			expectErr: false,
		},
		{
			name:      "正常値（負の数）",
			input:     -100,
			expected:  -100,
			expectErr: false,
		},
		{
			name:      "正常値（ゼロ）",
			input:     0,
			expected:  0,
			expectErr: false,
		},
		{
			name:      "正常値（int32最大値）",
			input:     math.MaxInt32,
			expected:  math.MaxInt32,
			expectErr: false,
		},
		{
			name:      "正常値（int32最小値）",
			input:     math.MinInt32,
			expected:  math.MinInt32,
			expectErr: false,
		},
		{
			name:      "エラー（int32最大値より大きい）",
			input:     math.MaxInt32 + 1,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "エラー（int32最小値より小さい）",
			input:     math.MinInt32 - 1,
			expected:  0,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SafeIntToInt32(tt.input)

			if tt.expectErr {
				if err == nil {
					t.Errorf("期待されたエラーが発生しませんでした")
				}
			} else {
				if err != nil {
					t.Errorf("予期しないエラーが発生しました: %v", err)
				}
				if result != tt.expected {
					t.Errorf("期待値 %d, 実際の値 %d", tt.expected, result)
				}
			}
		})
	}
}
