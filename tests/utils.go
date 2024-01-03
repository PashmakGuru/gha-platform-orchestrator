package tests

import (
	"io"
	"os"
)

func ReadFile(path string) (string, error) {
	content, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer content.Close()

	byteValue, err := io.ReadAll(content)

	return string(byteValue), err
}
