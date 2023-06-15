// Exercício 2 - Calcular média
// Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
// qual possamos passar N quantidade de números inteiros e devolva a média.
// Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro
package main

import "fmt"

func main() {
	fmt.Println(calcularMedia(4.0, 9.0, 10.0, 8.0))
}

func calcularMedia(notas ...float64) (media float64, err error) {
	soma := 0.0
	for i, nota := range notas {
		if notas[i] < 0.0 {
			return 0.0, fmt.Errorf("error: favor não informar números negativos.")
		}
		soma += nota
	}
	media = soma / float64(len(notas))
	return media, nil
}
