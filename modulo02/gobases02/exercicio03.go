// Exercício 3 - Calcular salário

// Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
// de horas trabalhadas por mês e na categoria do funcionário.
// Se a categoria for C, seu salário é de R$1.000 por hora
// Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h
// mensais
// Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h
// mensais

// Calcule o salário dos funcionários conforme as informações abaixo:
// - Funcionário de categoria C: 162h mensais
// - Funcionário de categoria B: 176h mensais
// - Funcionário de categoria A: 172h mensais

package main

import (
	"fmt"
)

func main(){
	fmt.Println("Funcionário de categoria C:", calcularSalario(10, "A"), "h")
	fmt.Println("Funcionário de categoria C:", calcularSalario(10, "A"), "h")
	fmt.Println("Funcionário de categoria C:", calcularSalario(10, "A"), "h")
}

func calcularSalario(horasTrabalhadas int, categoria string) float64{
	switch categoria{
	case "A":
		salario := 3000.00
		if horasTrabalhadas <= 160{
			return salario * float64(horasTrabalhadas)
		}
		return salario * float64(horasTrabalhadas) * 1.5

	case "B":
		salario := 1500.00
		if horasTrabalhadas <= 160{
			return salario * float64(horasTrabalhadas)
		}
		return salario * float64(horasTrabalhadas) * 1.2

	case "C":
		salario := 1000.00
		return salario * float64(horasTrabalhadas)

	default:
		fmt.Println("Opção incorreta! Favor inserir uma categoria válida.")
		return 0
	}
}