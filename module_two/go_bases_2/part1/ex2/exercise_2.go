package main

import (
	"errors"
	"fmt"
)

func main() {
	avg, err := calcAvg(-1, 20, 30, 40, 50, 60, 70, 80, 10, 10, 15, 1561, 5694, 1651, 65)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(avg)
}

func calcAvg(grades ...float64) (float64, error) {
	total := 0.0
	for _, value := range grades {
		if value < 0 {
			err := errors.New("Foi inserida um valor negativo!")
			return 0, err
		}
		total += value
	}
	return total / float64(len(grades)), nil
}
