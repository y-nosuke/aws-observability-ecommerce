package utils

import (
	"net/http"
)

func ContentType(image []byte) string {
	return http.DetectContentType(image[:512])
}

func Ext(image []byte) (contentType string, ext string) {
	contentType = http.DetectContentType(image[:512])
	switch contentType {
	case "image/png":
		ext = "png"
	case "image/jpeg":
		ext = "jpeg"
	case "image/gif":
		ext = "gif"
	case "image/webp":
		ext = "webp"
	default:
		ext = ""
	}
	return
}
