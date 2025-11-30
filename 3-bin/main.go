package main

import (
	"binjson/binscheme"
	"fmt"
)

func main() {
	binList := binscheme.NewBinList()
	id := "0001"
	private := false
	name := "first item"
	bin1 := binscheme.NewBin(id, private, name)
	binList.Bins = append(binList.Bins, *bin1)

	fmt.Println(binList)
}
