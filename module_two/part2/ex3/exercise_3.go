package main

import "fmt"

func main() {
	// Escolho resolver com pares de hashmap, assim tudo fica O(1)
	month, validMonth := selectMoth(16)

	if !validMonth {
		fmt.Println("O número de mês inseriro é inválido")
		return
	}

	fmt.Println("Mês de referência: ", month)
}

func selectMoth(month int) (string, bool) {
	months := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}
	value, ok := months[month]
	return value, ok
}
