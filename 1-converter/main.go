package main

import "fmt"

func main() {
	const USD_EUR = 0.86
	const USD_RUB = 81.0132
	const EUR_RUB = USD_RUB / USD_EUR
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
