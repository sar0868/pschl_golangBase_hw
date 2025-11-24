package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var currencies = []string{"USD", "EUR", "RUB"}

var ratesByUSD = map[string]float64{
	"USD": 1,
	"EUR": 0.86,
	"RUB": 81.0132,
}

func main() {

	fmt.Println("Конвертор валют.")
	for {
		var original, target string
		var count float64
		fmt.Print("Укажите исходную валюту (USD, EUR, RUB): ")
		original, err := InputData()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !checkInputCurrency(original) {
			continue
		}
		for {
			fmt.Print("Укажите количество исходной валюты: ")
			countStr, err := InputData()
			if err != nil {
				fmt.Println(err)
				continue
			}
			count, err = strconv.ParseFloat(countStr, 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if !checkCount(count) {
				continue
			}
			break
		}
		for {
			fmt.Printf("Укажите целевую валюту (%v): ", getCurrencyForExchange(original))
			target, err = InputData()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if !checkInputCurrency(target, original) {
				continue
			}
			break
		}
		CurrencyСalculation(count, original, target)
		break
	}
}

func CurrencyСalculation(count float64, currOriginal string, currTarget string) {
	result := calcRates(currOriginal, currTarget)
	fmt.Printf("%.2f\n", result*count)
}

func InputData() (string, error) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", fmt.Errorf("error: %w", err)
	}
	return input, nil
}

func checkInputCurrency(inputUser ...string) bool {
	if len(inputUser) == 2 && inputUser[0] == inputUser[1] && slices.Contains(currencies, inputUser[1]) {
		return false
	}
	return slices.Contains(currencies, inputUser[0])
}

func checkCount(count float64) bool {
	return count > 0
}

func getCurrencyForExchange(original string) string {
	var tempSlice []string
	for _, el := range currencies {
		if el == original {
			continue
		}
		tempSlice = append(tempSlice, el)
	}
	return fmt.Sprint(strings.Join(tempSlice, ", "))
}

func calcRates(original string, target string) float64 {
	originalByUSD := ratesByUSD[original]
	targetByUSD := ratesByUSD[target]
	return targetByUSD / originalByUSD
}
