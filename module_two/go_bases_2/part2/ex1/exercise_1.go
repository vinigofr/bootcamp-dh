package main

import "fmt"

type Aluno struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	RG        int    `json:"rg"`
	Data      string `json:"data_admissao"`
}

func (p Aluno) DetalhesAluno() {
	fmt.Println("=== INFORMAÇÕES DO ALUNO ===")
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Sobrenome:", p.Sobrenome)
	fmt.Println("RG:", p.RG)
	fmt.Println("Data de admissão:", p.Data)
}

func main() {
	var Alunos = []Aluno{
		{"Vinicius", "Gouveia", 000, "16/05/2022"},
		{"Raquel", "Gouveia", 000, "16/05/2022"},
		{"Bento", "Gouveia", 000, "16/05/2022"},
	}

	Alunos[0].DetalhesAluno()

}
