package main

import "fmt"

const (
	pequeno = "pequeno"
	medio   = "medio"
	grande  = "grande"
)

type produto struct {
	Tipo  string
	Nome  string
	Preco float64
}

type loja struct {
	produtos []Produto
}

func (l *loja) Adicionar(p Produto) {
	l.produtos = append(l.produtos, p)
}

func (l *loja) Total() float64 {
	total := 0.0
	for _, produto := range l.produtos {
		total += produto.CalcularCusto()
	}
	return total
}

type Produto interface {
	CalcularCusto() float64
}

func (p *produto) CalcularCusto() float64 {
	switch p.Tipo {
	case medio:
		return (p.Preco * 0.03) + p.Preco
	case grande:
		return (p.Preco * 0.06) + p.Preco + 2500
	default:
		return p.Preco
	}
}

type Ecommerce interface {
	Total() float64
	Adicionar(p Produto)
}

func novoProduto(tipo, nome string, preco float64) Produto {
	var newProduct = produto{
		tipo, nome, preco,
	}
	return &newProduct
}

func novaLoja() Ecommerce {
	return &loja{}
}

func main() {
	// Produtos que queremos adicionar
	var produtosParaAdicionar = []Produto{
		novoProduto("pequeno", "notebook", 2500), // 2500
		novoProduto("medio", "micro-ondas", 550), // 566.5
		novoProduto("grande", "geladeira", 4500),
	}

	// Instanciamos uma nova loja chamada Gouveia Empreendimentos
	var gouveiaEmpreendimentos = novaLoja()

	// Adicionamos os produtos
	for _, produto := range produtosParaAdicionar {
		gouveiaEmpreendimentos.Adicionar(produto)
	}

	// Exibimos seu custo
	fmt.Println(gouveiaEmpreendimentos.Total())
}
