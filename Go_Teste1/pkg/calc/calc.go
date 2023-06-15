package calc

func Subtracao(a, b int) int {
	return a - b
}

func SubtracaoI(a, b int) int {
	if a < b {
		return 0
	}
	return a - b
}

func BubbleSort(arrayNumber []int) []int {
	for i := 0; i < len(arrayNumber); i++ {
		for j := 0; j < len(arrayNumber)-1-i; j++ {
			if arrayNumber[j] > arrayNumber[j+1] {
				arrayNumber[j], arrayNumber[j+1] =  arrayNumber[j+1], arrayNumber[j]
			}
		}
	}
	return arrayNumber
}

func Divisao(numerador, denominador int) int{
	return numerador / denominador
}