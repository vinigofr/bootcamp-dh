package main

import (
	"fmt"
	"strings"
)

func main() {
	letterCount("Vinicius")
	withSlice("vinicius")
}

func letterCount(word string) {
	for _, value := range word {
		fmt.Println(string(value))
	}
}

func withSlice(word string) {
	splitedWord := strings.Split(word, "")

	for _, value := range splitedWord {
		fmt.Println(value)
	}
}
