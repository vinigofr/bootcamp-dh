package main

import "fmt"

type user struct {
	Nome      string
	Sobrenome string
	Email     string
	Senha     string
	Idade     int
}

type User interface {
	MudarNome(nome, sobrenome string)
	MudarIdade(idade int)
	MudarEmail(email string)
	MudarSenha(senha string)
}

func (u *user) MudarNome(nome, sobrenome string) {
	fmt.Printf("Nome alterado de %s %s ", u.Nome, u.Sobrenome)
	(*u).Nome = nome
	(*u).Sobrenome = sobrenome
	fmt.Println("para", (*u).Nome, (*u).Sobrenome)
}

//func (u *user) MudarIdade(idade int)    {}
//func (u *user) MudarEmail(email string) {}
//func (u *user) MudarSenha(senha string) {}

func main() {
	var novoUsuario user = user{Nome: "Vinicius", Sobrenome: "Gouveia", Email: "vini@gmail.com", Idade: 22, Senha: "123"}
	fmt.Println("Nome antigo", novoUsuario.Nome, novoUsuario.Sobrenome)
	novoUsuario.MudarNome("Silvio", "Santos")

}
