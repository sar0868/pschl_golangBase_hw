package main

import (
	"binjson/bins"
	"fmt"
)

func main() {
	binList := bins.NewBinList()
	id := "0001"
	private := false
	name := "first item"
	bin1 := bins.NewBin(id, private, name)
	binList.Bins = append(binList.Bins, *bin1)

	fmt.Println(binList)
}
