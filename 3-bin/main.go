package main

import (
	"binjson/bins"
	"fmt"
)

// Меню - создать бин, найти бины по id-private-name, посмотреть все бины, выход

var menu = map[string]func(*bins.BinList){
	"1": addBin,
	"2": findBins,
	"3": printBins,
}

func main() {
	bins := bins.NewBinList()

	fmt.Println("Bins")
Menu:
	for {
		choice := promptData(
			"1: add bin",
			"2: find bins",
			"3: print bins",
			"4: exit",
			"Input choice",
		)
		menuFunc := menu[choice]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(bins)
	}

	// binList := bins.NewBinList()
	// id := "0001"
	// private := false
	// name := "first item"
	// bin1 := bins.NewBin(id, private, name)
	// binList.Bins = append(binList.Bins, *bin1)

	// fmt.Println(binList)
}

func promptData(prompt ...string) string {
	for i, el := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", el)
		} else {
			fmt.Println(el)
		}
	}
	var result string
	fmt.Scan(&result)
	return result
}

func addBin(binsList *bins.BinList) {
	id := promptData("input id")
	inpPrivate := promptData("input private (1-true, 2 - false)")
	var private bool
	if inpPrivate == "1" {
		private = true
	} else {
		private = false
	}
	name := promptData("input name")
	newBin := bins.NewBin(id, private, name)
	binsList.Bins = append(binsList.Bins, *newBin)
}

func findBins(binList *bins.BinList) {}

func printBins(binList *bins.BinList) {
	for ind, bin := range binList.Bins {
		fmt.Println(ind, bin)
	}
}
