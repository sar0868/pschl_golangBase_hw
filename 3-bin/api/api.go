package api

import (
	"binjson/bins"
	"binjson/config"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Config config.Config

var option = map[string]func(*bins.BinListWithStorage){
	"1": FindId,
	"2": FindPrivate,
	"3": FindName,
}

func GetKey() {
	Config, err := config.Init()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(Config)
}

func PromptData(prompt ...string) string {
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

func AddBin(binsList *bins.BinListWithStorage) {
	for {
		id := PromptData("input id")
		inpPrivate := PromptData("input private (1-true, 2 - false)")
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
		if binsList.AddBin(*newBin) {
			return
		}
	}

}

func FindBins(binsList *bins.BinListWithStorage) {
	fmt.Println("\nselect an option to search:")
	choice := PromptData(
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

func FindId(binsList *bins.BinListWithStorage) {
	parameter := PromptData("specify id")
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
	PrintBins(&findBins)
}
func FindPrivate(binsList *bins.BinListWithStorage) {
	parameter := PromptData("specify private")
	result, ok := binsList.FindBins(parameter, func(bin bins.Bin, str string) bool {
		res, _ := strconv.ParseBool(str)
		return bin.Private == res
	})
	if !ok {
		fmt.Printf("Don't found bin for private: %v\n", parameter)
	}
	findBins := bins.BinListWithStorage{
		BinList: bins.BinList{
			Bins: result,
		},
	}
	PrintBins(&findBins)
}

func FindName(binsList *bins.BinListWithStorage) {
	parameter := PromptData("specify name")
	result, ok := binsList.FindBins(parameter, func(bin bins.Bin, str string) bool {
		return strings.Contains(bin.Name, str)
	})
	if !ok {
		fmt.Printf("Don't found bin for name: %v\n", parameter)
	}
	findBins := bins.BinListWithStorage{
		BinList: bins.BinList{
			Bins: result,
		},
	}
	PrintBins(&findBins)
}

func DeleteBinById(binsList *bins.BinListWithStorage) {
	id := PromptData("specify id")
	if !binsList.ContainsID(id) {
		fmt.Printf("Bins list does not contain bin with id=%v\n", id)
		return
	}
	if binsList.DeleteBins(id) {
		fmt.Printf("Delete bin with id=%v success\n", id)
	} else {
		fmt.Printf("Delete bin with id=%v failed\n", id)
	}
}

func PrintBins(binsList *bins.BinListWithStorage) {
	for _, bin := range binsList.Bins {
		fmt.Println("====")
		fmt.Println(bin.ToString())
	}
}
