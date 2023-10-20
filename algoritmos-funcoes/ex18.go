//Escreva uma função que receba um slice de inteiros e uma função como parâmetros. A função deve aplicar a função
//recebida em cada elemento do slice e retornar a soma de todos os resultados.
//Caso o slice esteja vazio, retorne um erro.

package main

import "fmt"

func main() {
	sliceInt := []int{50, 50, 50, 50}

	result, err := SumAllResults(sliceInt, MultiplyBy5)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}
}

func SumAllResults(sliceInt []int, f func([]int) []int) (int, error) {
	var sumAll int

	if len(sliceInt) == 0 {
		return 0, fmt.Errorf("slice cannot be empty")
	}

	for _, num := range f(sliceInt) {
		sumAll += num
	}

	return sumAll, nil
}

func MultiplyBy5(sliceInt []int) []int {
	var multipliedBy5 []int

	for _, num := range sliceInt {
		multipliedBy5 = append(multipliedBy5, num*5)
	}

	return multipliedBy5
}
