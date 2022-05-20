package main

import "fmt"

func main() {
	// 1. var sobrenome string = "Silva"
	/* Correta */

	// 2. var idade int = "25"
	/* Icorrreta */
	// var idade int = 25

	// 3. boolean := "false";
	/* Incorreta */
	// boolean := false

	// 4. var salario string = 4585.90
	/* Incorreta */
	// var salario float64 = 4585.90

	// 5. var nome string = "Fellipe"
	/* Correta */

	fmt.Println("// 1. var sobrenome string = \"Silva\"\n\t/* Correta */")
	fmt.Println("---")

	fmt.Printf("2. var idade int = \"25\"\n\t/* Icorrreta */\n\tvar idade int = 25")
	fmt.Println("---")

	fmt.Println("3. boolean := \"false\";\n\t/* Incorreta */\n\tboolean := false")
	fmt.Println("---")

	fmt.Println("4. var salario string = 4585.90\n\t/* Incorreta */\n\tvar salario float64 = 4585.90")
}
