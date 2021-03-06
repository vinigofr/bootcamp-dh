package main

import (
	"errors"
	"fmt"
)

var salario int

type customError error

var salarioMinimoErr customError
var impostoNecessarioErr customError

func main() {
	salarioMinimoErr = errors.New("erro: o salário digitado não está dentro do valor mínimo")
	impostoNecessarioErr = errors.New("Necessário pagamento de imposto")

	salario = 15001

	if salario < 15000 {
		fmt.Println(salarioMinimoErr.Error())
	} else {
		fmt.Println(impostoNecessarioErr.Error())
	}
}
