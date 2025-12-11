package main

import (
	"fmt"

	"binjson/api"
	"binjson/bins"
	"binjson/storage"
)

var menu = map[string]func(*bins.BinListWithStorage){
	"1": api.AddBin,
	"2": api.FindBins,
	"3": api.PrintBins,
	"4": api.DeleteBinById,
}

func main() {
	binsList := bins.NewBinList(storage.NewStorageJson("data.json"))

	fmt.Println("Bins")
Menu:
	for {
		fmt.Println("\nMenu:")
		choice := api.PromptData(
			"1: add bin",
			"2: find bins",
			"3: print bins",
			"4: delete bins by ID",
			"5: exit",
			"Input choice",
		)
		menuFunc := menu[choice]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(binsList)
	}
}
