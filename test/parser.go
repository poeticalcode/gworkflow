package test

import (
	"io"
	"os"
)

func XMLData(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	xmlData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return xmlData, nil
}
