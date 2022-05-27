package main

import (
	"fmt"
)

var salario int

type customError struct {
	msg string
}

func (ce customError) Error() string {
	return ce.msg
}

var salarioMinimoErr customError
var impostoNecessarioErr customError

func main() {
	salarioMinimoErr.msg = "erro: o salário digitado não está dentro do valor mínimo"
	impostoNecessarioErr.msg = "Necessário pagamento de imposto"

	salario = 15001

	if salario < 15000 {
		fmt.Println(salarioMinimoErr.Error())
	} else {
		fmt.Println(impostoNecessarioErr.Error())
	}
}
