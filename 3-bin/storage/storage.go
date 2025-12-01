package storage

import (
	"binjson/bins"
	"binjson/files"
	"encoding/json"
	"errors"
	"os"
)

func SaveBinToJson(bin bins.Bin, path string) error {
	data, err := json.Marshal(bin)
	if err != nil {
		return errors.New(err.Error())
	}
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

func GetBinListToJson(path string) (*bins.BinList, error) {
	data, err := files.Read(path)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	binList := bins.NewBinList()
	err = json.Unmarshal(data, binList)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return binList, nil
}
