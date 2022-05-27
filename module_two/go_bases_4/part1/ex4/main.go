package main

import (
	"errors"
	"fmt"
)

func CalcSalary(horasMes int, valorHora float64) (salario float64, err error) {
	salario = float64(horasMes) * valorHora

	if horasMes < 80 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}

	if salario >= 20000 {
		desconto := salario * 0.10
		salario -= desconto
	}
	return salario, nil
}

func main() {
	valor, err := CalcSalary(79, 17)

	if err != nil {
		err = fmt.Errorf("%s", err)
		fmt.Println(err)
		return
	}

	fmt.Println("Salário ->", valor)
}
