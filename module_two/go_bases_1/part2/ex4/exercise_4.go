package main

import "fmt"

func main() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	fmt.Println("A idade de Benjamin é", employees["Benjamin"])

	twentyOneYearsOld := 0
	for _, age := range employees {
		if age > 21 {
			twentyOneYearsOld++
		}
	}
	fmt.Println(twentyOneYearsOld, "funcionários tem 21 anos de idade")

	employees["Frederico"] = 25
	fmt.Println("Adicionado Frederico de 25 anos")

	delete(employees, "Pedro")
	fmt.Println("Pedro excluído")

}
