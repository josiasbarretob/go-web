// Exercício 3 - Declaração de variáveis
// Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
// Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
// declarar diferentes variáveis.
// Ajude o professor com as seguintes questões:
// 1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
// 2. Corrigir as incorrectas.
package main

import "fmt"

func main() {
	fmt.Println(`var 1nome string
		var sobrenome string
		var int idade
		1sobrenome := 6
		var licenca_para_dirigir = true
		var estatura da pessoa int
		quantidadeDeFilhos := 2`)
	fmt.Printf("CORREÇÃO:\n")
	fmt.Println(`var 1nome string -> var nome1 string
		var sobrenome string -> OK
		var int idade -> var idade int 
		1sobrenome := 6 -> sobrenome := 6
		var licenca_para_dirigir = true -> var licencaParaDirigir bool = true 
		var estatura da pessoa int -> var estaturaDaPessoa int
		quantidadeDeFilhos := 2 -> OK`)
}
