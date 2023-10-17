//Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com
//apenas os números pares contidos no slice. Caso o slice esteja vazio, retorne um erro.

package main

import "fmt"

func main() {
	sliceInt := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
	evenNumbers, err := FindEven(sliceInt)

	fmt.Printf("Antes: %d\n", sliceInt)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("Depois: ", evenNumbers)
	}
}

func FindEven(sliceInt []int) ([]int, error) {
	newSlice := []int{}

	for _, num := range sliceInt {
		if num%2 == 0 {
			newSlice = append(newSlice, num)
		}
	}

	if len(sliceInt) == 0 {
		return nil, fmt.Errorf("empty slice")
	}

	return newSlice, nil
}
