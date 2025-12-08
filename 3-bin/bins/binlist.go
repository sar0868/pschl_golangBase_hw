package bins

type Storage interface {
	SaveBinsList(BinList) error
	GetBinsList() (*BinList, error)
}
type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBinList(storage Storage) *BinListWithStorage {
	return &BinListWithStorage{
		BinList: BinList{},
		storage: storage,
	}
}

type BinListWithStorage struct {
	BinList
	storage Storage
}
