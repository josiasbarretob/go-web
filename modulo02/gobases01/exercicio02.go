// Exercício 2 - Clima
// Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e
// pressão atmosférica de diferentes lugares.
// 1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem
// possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
// 2. Imprima os valores no console.
// 3. Quais tipos de dados serão atribuídos a essas variáveis?
package main

import "fmt"

var (
	temperatura        int    = 21
	umidade            string = "80%"
	pressaoAtmosferica string = "3 ATM"
)

func main() {
	fmt.Printf("Temperatura: %v - Umidade: %v - Pressão Atmosférica: %v\n", temperatura, umidade, pressaoAtmosferica)
}
