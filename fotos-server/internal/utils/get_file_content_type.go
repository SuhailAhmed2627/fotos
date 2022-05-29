package utils

import (
	"mime/multipart"
	"net/http"
)

func GetFileContentType(out multipart.File) (string, error) {
	defer out.Close()
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	ctype := http.DetectContentType(buffer)
	return ctype, nil
}
