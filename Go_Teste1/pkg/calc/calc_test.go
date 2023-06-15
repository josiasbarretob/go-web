package calc_test

import (
	"go_teste1/calc/pkg/calc"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtracao(t *testing.T) {
	number1 := 100
	number2 := 70
	expectedResult := 30

	//Execução do Teste
	resultAtual := calc.Subtracao(number1, number2)
	if resultAtual != expectedResult {
		t.Errorf("A funcao subtracao() retornou o resultado %v, mas esperava %v!", resultAtual, expectedResult)
	}
}

func TestSubtracaoI(t *testing.T) {
	expectedResult := 0
	resultAtual := calc.SubtracaoI(5, 5)

	// if resultAtual != expectedResult {
	// 	t.Errorf("A funcao subtracao() retornou o resultado %v, mas esperava %v!", resultAtual, expectedResult)
	// }
	assert.Equal(t, expectedResult, resultAtual, "o resultado atual precisa ser igual ao resultado esperado")
}

func TestBubbleSort(t *testing.T) {
	numberArray := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	expectedResult := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actualResult := calc.BubbleSort(numberArray)

	// assert.Equal(t, expectedResult, actualResult, "o resultado atual precisa ser igual ao resultado esperad")
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Array ordenado incorretamente. Esperado: %v, Obtido: %v", expectedResult	, actualResult)
	}
}

func TestDivisao(t *testing.T){
	actualResult := calc.Divisao(10, 2)

	assert.NotNil(t, actualResult)
}