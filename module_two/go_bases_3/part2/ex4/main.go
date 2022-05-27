package main

import (
	"fmt"
	"math/rand"
)

func insertionSort(items []int) []int {
	var length = len(items)

	for i := 1; i < length; i++ {
		currentPosition := i

		for currentPosition > 0 {
			if items[currentPosition-1] > items[currentPosition] {
				items[currentPosition-1], items[currentPosition] = items[currentPosition], items[currentPosition-1]
			}
			currentPosition -= 1
		}
	}
	return items
}

func main() {
	var1 := rand.Perm(10000)

	fmt.Println(var1)
	fmt.Println(insertionSort(var1))
}
