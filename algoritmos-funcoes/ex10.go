//Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com os
//valores ordenados de forma crescente. Caso o slice esteja vazio, retorne um erro.

package main

import "fmt"

func main() {
	sliceInt := []int{100, 72, 1, 12, 29, 21}

	fmt.Println("Slice original: ", sliceInt)

	ordenateSlice, err := OrdenateNumbers(sliceInt)
	if err == nil {
		fmt.Print("Slice ordenado: ", ordenateSlice)
	} else {
		fmt.Print(err)
	}
}

func OrdenateNumbers(sliceInt []int) ([]int, error) {
	if len(sliceInt) == 0 {
		return nil, fmt.Errorf("empty slice")
	}

	for i := 0; i < len(sliceInt); i++ {
		for j := 0; j < len(sliceInt); j++ {
			if sliceInt[i] < sliceInt[j] {
				sliceInt[i], sliceInt[j] = sliceInt[j], sliceInt[i]
			}
		}
	}

	return sliceInt, nil
}
