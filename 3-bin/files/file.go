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

func Write(data []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New(err.Error())
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		return errors.New("error write data")
	}
	return nil
}
