package main

import (
	"errors"
	"testing"
)

func TestCalcSalary(t *testing.T) {

	verificaErr := func(t *testing.T, errObtido, errEsperado error) {
		t.Helper()
		if errObtido.Error() != errEsperado.Error() {
			t.Errorf("O resultado '%v' é diferente de '%v'", errObtido, errEsperado)
		}
	}

	verificaResultado := func(t *testing.T, resultadoObtido, resultadoEsperado float64) {
		t.Helper()
		if resultadoObtido != resultadoEsperado {
			t.Errorf("O resultado '%v' é diferente de '%v'", resultadoObtido, resultadoEsperado)
		}
	}

	verificaNil := func(t *testing.T, nilVar interface{}) {
		t.Helper()
		if nilVar != nil {
			t.Errorf("Foi esperado que não houvesse um erro, porém recebeu '%v'", nilVar)
		}
	}

	t.Run("Testa com numero de horas abaixo do aceito", func(t *testing.T) {
		resultado, err := CalcSalary(50, 75.00)
		resultadoEsperado := 0.0
		erroEsperado := errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")

		verificaResultado(t, resultado, resultadoEsperado)
		verificaErr(t, err, erroEsperado)
	})

	t.Run("Testa se funciona ao passar de 80 horas", func(t *testing.T) {
		resultado, err := CalcSalary(80, 75.00)
		resultadoEsperado := 6000.00

		verificaResultado(t, resultado, resultadoEsperado)
		verificaNil(t, err)
	})

	t.Run("Testa se o cálculo de imposto subtrai 10% ao passar de um salário de 20000", func(t *testing.T) {
		resultado, err := CalcSalary(1500, 1000)
		resultadoEsperado := 1350000.00

		verificaResultado(t, resultado, resultadoEsperado)
		verificaNil(t, err)
	})
}
