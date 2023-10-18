//Escreva uma função que receba dois slices de inteiros como parâmetros e retorne um novo slice contendo apenas os
//números presentes nos dois slices. Caso um dos slices esteja vazio, retorne um erro.

package main

import "fmt"

func main() {
	sliceX := []int{1, 2, 3, 16, 9, 17, 11, 15, 14, 6}
	sliceY := []int{13, 9, 16, 5, 11, 12, 18, 19, 14, 3}

	newSlice, err := FindRepeatedNumbers(sliceX, sliceY)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(newSlice)
	}
}

func FindRepeatedNumbers(sliceX, sliceY []int) ([]int, error) {
	var newSlice []int

	if len(sliceX) == 0 || len(sliceY) == 0 {
		return nil, fmt.Errorf("any slice cannot be empty")
	}

	for _, numX := range sliceX {
		for _, numY := range sliceY {
			if numX == numY {
				newSlice = append(newSlice, numY)
			}
		}
	}

	return newSlice, nil
}
