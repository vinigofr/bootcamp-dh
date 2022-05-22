package main

import "fmt"

type AlunoInfo struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	RG        int    `json:"rg"`
	Data      string `json:"data_admissao"`
}

type Alunos interface {
	Detalhes()
}

func (aluno AlunoInfo) Detalhes() {
	fmt.Println("=== INFORMAÇÕES DO ALUNO ===")
	fmt.Println("Nome:", (aluno).Nome)
	fmt.Println("Sobrenome:", (aluno).Sobrenome)
	fmt.Println("RG:", (aluno).RG)
	fmt.Println("Data de admissão:", (aluno).Data)
}

func Detalhes(a Alunos) {
	a.Detalhes()
}

func main() {
	var Alunos = []AlunoInfo{{"Vinicius", "Gouveia", 000, "16/05/2022"}}

	Alunos[0].Detalhes()

	Detalhes(Alunos[0])

}
