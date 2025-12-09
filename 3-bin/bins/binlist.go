package bins

import "fmt"

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

func (bins *BinListWithStorage) FindBins(parameter string, checker func(bin Bin, str string) bool) ([]Bin, bool) {
	var result []Bin
	for _, bin := range bins.Bins {
		if checker(bin, parameter) {
			result = append(result, bin)
		}
	}
	return result, len(result) != 0
}

func NewBinList(storage Storage) *BinListWithStorage {
	binsList := &BinListWithStorage{
		BinList: BinList{},
		storage: storage,
	}
	list, err := storage.GetBinsList()
	if err != nil {
		fmt.Printf("error get bins list: %v", err)
		return &BinListWithStorage{
			BinList: BinList{},
			storage: storage,
		}
	}
	binsList.BinList = *list
	return binsList
}

func (bins *BinListWithStorage) AddBin(bin Bin) {
	bins.Bins = append(bins.Bins, bin)
	err := bins.save()
	if err != nil {
		fmt.Println(err)
	}
}

func (bins *BinListWithStorage) save() error {
	err := bins.storage.SaveBinsList(bins.BinList)
	if err != nil {
		return fmt.Errorf("error save bin list: %w", err)
	}
	return nil
}
