package main

import "fmt"

func main() {
	result1 := salaryCalc('C', 162)
	result2 := salaryCalc('B', 176)
	result3 := salaryCalc('A', 172)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func salaryCalc(category rune, workedHours int) float64 {
	// Map com salários e seus tipos
	salaryMap := map[rune]float64{
		'C': 1000.00,
		'B': 1500.00,
		'A': 3000.00,
	}

	// Calculo dinâmico de salário * horas
	calculatedSalary := salaryMap[category] * float64(workedHours)

	if category == 'B' && workedHours > 160 {
		calculatedSalary = (calculatedSalary * 0.2) + calculatedSalary
	}
	if category == 'A' && workedHours > 160 {
		calculatedSalary = (calculatedSalary * 0.5) + calculatedSalary
	}
	return calculatedSalary
}
