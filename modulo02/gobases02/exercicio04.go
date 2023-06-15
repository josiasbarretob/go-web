// Exercício 4 - Cálculo de estatísticas

// Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
// notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
// de suas notas.
// Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
// máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
// definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
// indicado na função anterior
package main

import "fmt"

func main() {
	// fmt.Println(calcularMinimo(9, 8, 7, 6, 5, 4, 3, 2, 1))
	// fmt.Println(calcularMinimo(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))

	// fmt.Println(calcularMaximo(9, 8, 7, 6, 5, 4, 3, 2, 1))
	// fmt.Println(calcularMaximo(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))

	// fmt.Println(calcularMedia(9, 8, 7, 6, 5, 4, 3, 2, 1))
	// fmt.Println(calcularMedia(4, 6, 10))
	funcaoMinimo := verificarTipoCalculo("minimo")
	funcaoMaximo := verificarTipoCalculo("maximo")
	funcaoMedia := verificarTipoCalculo("media")

	min := funcaoMinimo(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	max := funcaoMaximo(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	med := funcaoMedia(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Printf("Valor Mínimo: %.2f - Valor Máximo: %.2f - Média dos Valores: %.2f\n", min, max, med)
}

func verificarTipoCalculo(tipo string) (func (valor... float64) float64){
	switch tipo{
	case "minimo":
		return calcularMinimo
	case "maximo":
		return calcularMaximo
	case "media":
		return calcularMedia
	default:
		fmt.Println("Favor informar um tipo de cálculo válido! Ex. minimo, maximo e média.")
	}
	return nil
}

func calcularMinimo(notas ...float64) float64 {
	minimo := notas[0]
	for i := 0; i < len(notas); i++ {
		if notas[i] < minimo {
			minimo = notas[i]
		}
	}
	return minimo
}

func calcularMaximo(notas ...float64) float64 {
	maximo := notas[0]
	for i := 0; i < len(notas); i++ {
		if notas[i] > maximo {
			maximo = notas[i]
		}
	}
	return maximo
}

func calcularMedia(notas ...float64) float64 {
	soma := 0.0
	for _, nota := range notas {
		soma += nota
	}
	return soma / float64(len(notas))
}
