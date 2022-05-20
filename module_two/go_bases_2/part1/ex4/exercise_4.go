package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, _ := chooseCalcMethod(minimum)
	avgFunc, _ := chooseCalcMethod(average)
	maxFunc, _ := chooseCalcMethod(maximum)

	result1 := minFunc(1, 2, 3, 45, 6, 7, 878, 78)
	result2 := avgFunc(1, 2, 3, 45, 6, 7, 878, 78)
	result3 := maxFunc(1, 2, 3, 45, 6, 7, 878, 78)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func getMinimumValue(values ...int) int {
	currentMinimumValue := values[0]

	for _, value := range values {
		if currentMinimumValue > value {
			currentMinimumValue = value
		}
	}

	return currentMinimumValue
}

func getMaxValue(values ...int) int {
	currentMaxValue := values[0]

	for _, value := range values {
		if currentMaxValue < value {
			currentMaxValue = value
		}
	}

	return currentMaxValue
}

func getAverageValue(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}

	return total / len(values)
}

func chooseCalcMethod(method string) (func(values ...int) int, error) {
	// Referência: https://stackoverflow.com/questions/52298469/map-to-store-generic-type-functions-in-go
	functionMap := map[string]func(values ...int) int{
		minimum: getMinimumValue,
		average: getAverageValue,
		maximum: getMaxValue,
	}

	selectedFunction, ok := functionMap[method]

	if !ok {
		return nil, errors.New("Não foi possível localizar a função desejada!")
	}
	return selectedFunction, nil
}
