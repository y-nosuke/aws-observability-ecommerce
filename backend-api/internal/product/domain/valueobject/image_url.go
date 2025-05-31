package valueobject

import (
	"fmt"
	"strings"
)

// ImageURL は商品画像のURLを表す値オブジェクト
type ImageURL struct {
	value string
}

// NewImageURL は新しいImageURLを作成する
func NewImageURL(value string) (*ImageURL, error) {
	if value == "" {
		return nil, fmt.Errorf("image URL cannot be empty")
	}

	// URLの形式チェック
	if !strings.HasPrefix(value, "http://") && !strings.HasPrefix(value, "https://") {
		return nil, fmt.Errorf("invalid image URL format")
	}

	return &ImageURL{
		value: value,
	}, nil
}

// Value はURLの値を取得する
func (u *ImageURL) Value() string {
	return u.value
}

// String は文字列表現を返す
func (u *ImageURL) String() string {
	return u.value
}
