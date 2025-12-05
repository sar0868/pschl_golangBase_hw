package files

import (
	"errors"
	"os"
	"strings"
)

func Read(path string) ([]byte, error) {
	if !checkTypeFile(path) {
		return nil, errors.New("extension not json")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func checkTypeFile(filename string) bool {
	arr := strings.Split(filename, ".")
	if len(arr) < 2 {
		return false
	}
	if arr[len(arr)-1] != "json" {
		return false
	}
	return true
}
