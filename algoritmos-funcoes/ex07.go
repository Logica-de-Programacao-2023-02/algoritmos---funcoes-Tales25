//Crie uma função que receba um slice de inteiros e uma função como parâmetros. A função deve aplicar a função recebida
//em cada elemento do slice e retornar um novo slice com os resultados. Caso o slice esteja vazio, retorne um erro.

package main

import "fmt"

func main() {
	sliceInt := []int{2, 4, 6, 8, 10}

	resultDouble, err := Operation(sliceInt, Double)
	resultDivide, err := Operation(sliceInt, Divide)

	fmt.Println("Original: ", sliceInt)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Dobrado: ", resultDouble)
		fmt.Println("Metade: ", resultDivide)
	}
}

func Double(sliceInt []int) []int {
	var doubled []int

	for _, num := range sliceInt {
		doubled = append(doubled, num*2)
	}

	return doubled
}

func Divide(sliceInt []int) []int {
	var divided []int

	for _, num := range sliceInt {
		divided = append(divided, num/2)
	}

	return divided
}

func Operation(sliceInt []int, f func([]int) []int) ([]int, error) {
	if len(sliceInt) == 0 {
		return nil, fmt.Errorf("empty slice")
	}

	return f(sliceInt), nil
}
