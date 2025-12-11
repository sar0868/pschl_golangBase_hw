package bins

import (
	"fmt"
	"slices"
	"strings"
)

type Storage interface {
	SaveBinsList(BinList) error
	GetBinsList() (*BinList, error)
}
type BinList struct {
	Bins []Bin `json:"bins"`
}

type BinListWithStorage struct {
	BinList
	storage Storage
}

func (bins *BinListWithStorage) DeleteBins(id string) bool {
	temp := len(bins.Bins)
	bins.Bins = slices.DeleteFunc(bins.Bins, func(bin Bin) bool {
		return strings.Contains(bin.Id, id)
	})
	if temp > len(bins.Bins) {
		if err := bins.save(); err != nil {
			fmt.Println(err)
		}
		return true
	}
	return false
}

func (bins *BinListWithStorage) FindBins(parameter string, checker func(bin Bin, str string) bool) ([]Bin, bool) {
	var result []Bin
	for _, bin := range bins.Bins {
		if checker(bin, parameter) {
			result = append(result, bin)
		}
	}
	return result, len(result) != 0
}

func NewBinList(storage Storage) (*BinListWithStorage, error) {
	binsList := &BinListWithStorage{
		BinList: BinList{},
		storage: storage,
	}
	list, err := storage.GetBinsList()
	if err != nil {
		fmt.Printf("error get bins list: %v", err)
		return nil, fmt.Errorf("error: %w", err)
	}
	binsList.BinList = *list
	return binsList, nil
}

func (bins *BinListWithStorage) AddBin(bin Bin) bool {
	if bins.ContainsID(bin.Id) {
		fmt.Printf("Bin with id=%v already exists\n", bin.Id)
		return false
	}
	bins.Bins = append(bins.Bins, bin)
	err := bins.save()
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func (bins *BinListWithStorage) save() error {
	err := bins.storage.SaveBinsList(bins.BinList)
	if err != nil {
		return fmt.Errorf("error save bin list: %w", err)
	}
	return nil
}

func (bins *BinListWithStorage) ContainsID(id string) bool {
	_, ok := bins.FindBins(id, func(bin Bin, str string) bool {
		return bin.Id == str
	})
	return ok
}
