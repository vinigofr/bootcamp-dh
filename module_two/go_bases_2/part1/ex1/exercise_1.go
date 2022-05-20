package main

import "fmt"

func main() {
	fmt.Println(taxSalary(150000))
}

func taxSalary(salary float64) float64 {
	if salary < 150000 {
		return salary * 0.17
	}
	return salary * 0.27
}
