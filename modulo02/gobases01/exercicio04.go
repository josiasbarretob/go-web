// Exercício 4 - Tipos de dados
// Um estudante de programação tentou fazer declarações de variáveis com seus respectivos
// tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e
// pediu a ajuda de um desenvolvedor experiente que pode:
// ● Revisar o código e realizar as devidas correções.
package main

import "fmt"

func main() {
	fmt.Printf(`
	var sobrenome string = "Silva"
	var idade int = "25"
	boolean := "false";
	var salario string = 4585.90
	var nome string = "Fellipe"
	`)
	fmt.Printf("CORREÇÃO:\n")
	fmt.Printf(`
	var sobrenome string = "Silva" -> OK
	var idade int = "25" -> var idade int = 25
	boolean := "false"; -> boolean := false
	var salario string = 4585.90 -> var salario float = 4585.90
	var nome string = "Fellipe" -> OK
	`)
}
