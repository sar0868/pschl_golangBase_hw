package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var currencies = []string{"USD", "EUR", "RUB"}

func main() {
	const USD_EUR = 0.86
	const USD_RUB = 81.0132
	const EUR_RUB = USD_RUB / USD_EUR

	fmt.Println("Конвертор валют.")
	for {
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
			count, err := strconv.Atoi(countStr)
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
			target, err := InputData()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if !checkInputCurrency(target, original) {
				continue
			}
			break
		}
		fmt.Println("OK")
		break
	}
	// for {
	// 	menu()
	//
	// 	switch choice {
	// 	case "1":
	// 		if checkInputCurrency()
	// 		menuStep2()
	// 	case "2":
	// 		return
	// 	default:
	// 		continue
	// 	}

	// }
}

func CurrencyСalculation(count float64, currOriginal string, currTarget string) {

}

func InputData() (string, error) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", fmt.Errorf("error: %w", err)
	}
	return input, nil
}

func menu() {

	fmt.Println("Выберете вариант")

	fmt.Println("2: Выход")
}

func checkInputCurrency(inputUser ...string) bool {
	if len(inputUser) == 2 && inputUser[0] == inputUser[1] && slices.Contains(currencies, inputUser[1]) {
		return false
	}
	return slices.Contains(currencies, inputUser[0])
}

func checkCount(count int) bool {
	return count > 0
}

func getCurrencyForExchange(original string) string {
	indOriginal := slices.Index(currencies, original)
	tempSlice := currencies[:indOriginal]
	// if indOriginal == len(currencies)-1 {
	// 	tempSlice = currencies[:indOriginal]
	// } else {
	tempSlice = append(tempSlice, currencies[indOriginal+1:]...)
	// }
	return fmt.Sprint(strings.Join(tempSlice, ", "))
}
