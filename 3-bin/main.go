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
	var choice string
	fmt.Scan(&choice)
	return choice
}

func addBin(bins *bins.BinList) {}

func findBins(bins *bins.BinList) {}

func printBins(bins *bins.BinList) {}
