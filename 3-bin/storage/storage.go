package storage

import (
	"binjson/bins"
	"binjson/files"
	"encoding/json"
	"errors"
)

type StorageJson struct {
	filename string
}

// GetBinsList implements bins.Storage.

func NewStorageJson(filename string) *StorageJson {
	return &StorageJson{
		filename: filename,
	}
}

func (jsonStorage *StorageJson) SaveBinsList(bins bins.BinList) error {
	data, err := json.Marshal(bins)
	if err != nil {
		return errors.New(err.Error())
	}
	err = files.Write(data, jsonStorage.filename)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (jsonStorage *StorageJson) GetBinsList() (*bins.BinList, error) {
	data, err := files.Read(jsonStorage.filename)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	binList := bins.NewBinList(NewStorageJson(jsonStorage.filename))
	err = json.Unmarshal(data, binList)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &binList.BinList, nil
}
