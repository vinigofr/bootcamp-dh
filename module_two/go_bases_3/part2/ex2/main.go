package main

import "fmt"

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Usuario struct {
	Nome     string
	Produtos []Produto
}

// novoProduto recebe nome e preço como parâmetro e retorna um novo Produto
func novoProduto(nome string, preco float64) Produto {
	return Produto{
		Nome:       nome,
		Preco:      preco,
		Quantidade: 1,
	}
}

// adicionarProduto adiciona um respectivo produto a um usuário
func adicionarProduto(usr *Usuario, pdt Produto) {
	usr.Produtos = append(usr.Produtos, pdt)
	fmt.Println("Produto adicionado com sucesso ao usuário", usr.Nome)
}

// deletarProdutos apaga todos os produtos que um usuário possui
// Limpando um slice: https://yourbasic.org/golang/clear-slice/
func deletarProdutos(usr *Usuario) {
	usr.Produtos = nil
	fmt.Println("Produtos de", usr.Nome, "eliminados da base de dados")
}

func main() {
	computador := novoProduto("Computador", 5000)
	usuarioVinicius := Usuario{Nome: "Vini"}

	adicionarProduto(&usuarioVinicius, computador)
	fmt.Println(usuarioVinicius.Produtos[0])

	deletarProdutos(&usuarioVinicius)
}
