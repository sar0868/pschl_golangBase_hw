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

func NewBinList(storage Storage) *BinListWithStorage {
	list, err := storage.GetBinsList()
	if err != nil {
		fmt.Printf("error get bins list: %v", err)
		return &BinListWithStorage{
			BinList: BinList{},
			storage: storage,
		}
	}
	return &BinListWithStorage{
		BinList: *list,
		storage: storage,
	}
	// return &BinListWithStorage{
	// 	BinList: BinList{},
	// 	storage: storage,
	// }
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
