package utils

import (
	"fmt"
	"math"
)

// SafeIntToInt32 は int を int32 に安全に変換します
// int32 の範囲外の場合はエラーを返します
func SafeIntToInt32(value int) (int32, error) {
	if value < math.MinInt32 || value > math.MaxInt32 {
		return 0, fmt.Errorf("value %d is out of int32 range [%d, %d]", value, math.MinInt32, math.MaxInt32)
	}
	return int32(value), nil
}
