package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Handler func([]float64) float64

var operations = map[string]Handler{
	"SUM": sum,
	"AVG": avg,
	"MED": med,
}

func main() {
	fmt.Println("Калькулятор")
Menu:
	for {
		menu()
		var inputUser string
		_, err := fmt.Scan(&inputUser)
		if err != nil {
			fmt.Println(err)
			continue
		}
		funOperation, ok := operations[inputUser]
		if !ok {
			break Menu
		}

		array, err := inputData()
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := funOperation(array)
		fmt.Printf("%.2f\n", result)
	}
}

func menu() {
	fmt.Println("\nВыберете операцию:")
	fmt.Println("Сложение - введите SUM")
	fmt.Println("Средняя - введите AVG")
	fmt.Println("Медина - введите MED")
	fmt.Println("Для выхода введите 0")
}

func inputData() ([]float64, error) {
	result := []float64{}
	fmt.Println("Введите числа через запятую")
	reader := bufio.NewReader(os.Stdin)

	arrInput, err := reader.ReadString('\n')

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	arr := strings.Split(arrInput, ",")
	for _, el := range arr {
		el = strings.TrimSpace(el)
		num, err := strconv.ParseFloat(el, 64)
		if err != nil {
			return nil, fmt.Errorf("error: %w", err)
		}
		result = append(result, num)
	}
	return result, nil
}

func sum(arr []float64) float64 {
	result := 0.0
	for _, el := range arr {
		result += el
	}
	return result
}

func avg(arr []float64) float64 {
	sum := sum(arr)
	count := len(arr)
	return sum / float64(count)
}

func med(arr []float64) float64 {
	sort.Float64s(arr)
	if len(arr)%2 != 0 {
		return arr[(len(arr)-1)/2]
	}
	ind1 := (len(arr) / 2) - 1
	ind2 := ind1 + 1
	return avg([]float64{arr[ind1], arr[ind2]})
}
