package main

import (
	"binjson/bins"
	"binjson/storage"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Меню - создать бин, найти бины по id-private-name, посмотреть все бины, выход

var menu = map[string]func(*bins.BinListWithStorage){
	"1": addBin,
	"2": findBins,
	"3": printBins,
	"4": deleteBinById,
}

var option = map[string]func(*bins.BinListWithStorage){
	"1": findId,
	"2": findPrivate,
	"3": findName,
}

func main() {
	binsList := bins.NewBinList(storage.NewStorageJson("data.json"))

	fmt.Println("Bins")
Menu:
	for {
		choice := promptData(
			"\n1: add bin",
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

func addBin(binsList *bins.BinListWithStorage) {
	id := promptData("input id")
	inpPrivate := promptData("input private (1-true, 2 - false)")
	var private bool
	if inpPrivate == "1" {
		private = true
	} else {
		private = false
	}
	var name string
	for {
		fmt.Println("input name (after you finish entering your name, press Enter)")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			name = scanner.Text()
			name = strings.TrimSpace(name)
		}
		if name != "" {
			break
		}
	}

	newBin := bins.NewBin(id, private, name)
	binsList.AddBin(*newBin)
}

func findBins(binsList *bins.BinListWithStorage) {
	fmt.Println("\nselect an option to search:")
	choice := promptData(
		"1: id",
		"2: private",
		"3: name",
		"4: return to menu",
		"Input choice",
	)
	menuFind := option[choice]
	if menuFind == nil {
		return
	}
	menuFind(binsList)
}

func findId(binsList *bins.BinListWithStorage) {
	parameter := promptData("specify id")
	result, ok := binsList.FindBins(parameter, func(bin bins.Bin, str string) bool {
		return bin.Id == str
	})
	if !ok {
		fmt.Printf("Don't found bin for id: %v", parameter)
	}
	findBins := bins.BinListWithStorage{
		BinList: bins.BinList{
			Bins: result,
		},
	}
	printBins(&findBins)
}
func findPrivate(binsList *bins.BinListWithStorage) {
	parameter := promptData("specify private")
	result, ok := binsList.FindBins(parameter, func(bin bins.Bin, str string) bool {
		res, _ := strconv.ParseBool(str)
		return bin.Private == res
	})
	if !ok {
		fmt.Printf("Don't found bin for private: %v", parameter)
	}
	findBins := bins.BinListWithStorage{
		BinList: bins.BinList{
			Bins: result,
		},
	}
	printBins(&findBins)
}

func findName(binsList *bins.BinListWithStorage) {
	parameter := promptData("specify name")
	result, ok := binsList.FindBins(parameter, func(bin bins.Bin, str string) bool {
		return strings.Contains(bin.Name, str)
	})
	if !ok {
		fmt.Printf("Don't found bin for name: %v", parameter)
	}
	findBins := bins.BinListWithStorage{
		BinList: bins.BinList{
			Bins: result,
		},
	}
	printBins(&findBins)
}

func deleteBinById(binsList *bins.BinListWithStorage) {
	//
}

func printBins(binsList *bins.BinListWithStorage) {
	for _, bin := range binsList.Bins {
		fmt.Println("====")
		fmt.Println(bin.ToString())
	}
}
