package bins

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBinList() *BinList {
	return &BinList{}
}
