//Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.

package main

import "fmt"

func main() {
	sliceInt := []int{17, 49, 21, 56, 1, 18, 100, 60, 24, 52, 37, 94, 39, 10, 81, 84, 47, 69, 41, 67}
	secondMaxValue := secondMaxValue(sliceInt)
	fmt.Print(secondMaxValue)
}

func secondMaxValue(sliceInt []int) int {
	for i := 0; i < len(sliceInt); i++ {
		for j := 0; j < len(sliceInt); j++ {
			//I > J (decrescente); I < J (crescente)
			if sliceInt[i] > sliceInt[j] {
				sliceInt[i], sliceInt[j] = sliceInt[j], sliceInt[i]
			}
		}
	}
	//segundo maior na ordem decrescente é o segundo termo, index 1
	result := sliceInt[1]
	return result
}
