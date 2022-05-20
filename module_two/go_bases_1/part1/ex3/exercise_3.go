package main

import "fmt"

func main() {
	// 1. var 1nome string
	/* Correção: Está incorreta pois está começando com número no nome da var */
	// var nome1 string

	// 2. var sobrenome string
	/* Correção: OK */

	// 3. var int idade
	/* Correção: O tipo da variável está vindo antes de seu identificador. */
	// var idade int

	// 4. 1sobrenome := 6
	/* Correção: Está incorreta pois está começando com número no nome da var */
	/* Correção: Declaração curta está fora do escopo de uma função */
	//func() {
	//	sobrenome1 := 6
	//	fmt.Println(sobrenome1)
	//}()

	// 5. var licenca_para_dirigir = true
	/* Correção: Está correta a sintaxe, porém foge dos padrões camelCase */
	//var licencaParaDirigir = true

	// 6. var estatura da pessoa int
	/* Correção: Está incorreta a sintaxe, contém espaços e foge dos padrões camelCase */
	//var estaturaDaPessoa int

	// 7. quantidadeDeFilhos := 2
	/* Correção: OK, se considerarmos esta variável estar dentro de uma função */

	fmt.Println("EXERCÍCIO 3")

	fmt.Printf("1. var 1nome string\n" +
		"/* Correção: Está incorreta pois está começando com número no nome da var */\n" +
		"var nome1 string\n===\n")

	fmt.Printf("2. var sobrenome string\n" +
		"/* Correção: OK */\n===\n")

	fmt.Printf("3. var int idade\n" +
		"/* Correção: O tipo da variável está vindo antes de seu identificador. */\n" +
		"var idade int\n===\n")

	fmt.Printf("4. 1sobrenome := 6\n" +
		"/* Correção: Está incorreta pois está começando com número no nome da var */\n" +
		"/* Correção: Declaração curta está fora do escopo de uma função */\n" +
		"func() {\n\tsobrenome1 := 6\n\tfmt.Println(sobrenome1)\n\t}()\n===\n")

	fmt.Printf("5. var licenca_para_dirigir = true\n" +
		"/* Correção: Está correta a sintaxe, porém foge dos padrões camelCase */\n" +
		"var licencaParaDirigir = true\n===\n")

	fmt.Printf("6. var estatura da pessoa int\n" +
		"/* Correção: Está incorreta a sintaxe, contém espaços e foge dos padrões camelCase */\n" +
		"var estaturaDaPessoa int\n===\n")

	fmt.Printf("7. quantidadeDeFilhos := 2\n" +
		"/* Correção: OK, se considerarmos esta variável estar dentro de uma função */\n\n")
}
