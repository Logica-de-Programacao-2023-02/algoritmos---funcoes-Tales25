//Crie uma função que receba um slice de inteiros e um valor inteiro, e retorne a posição do primeiro
//elemento igual ao valor no slice. Caso não encontre, retorne -1.

package main

import "fmt"

func main() {
	sliceInt := []int{17, 49, 21, 56, 1, 18, 100, 60, 24, 52, 37, 94, 39, 10, 81, 84, 47, 69, 41, 67}
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	fmt.Print(findNumber(sliceInt, num))
}

func findNumber(sliceInt []int, int int) int {
	for i := 0; i < len(sliceInt); i++ {
		if sliceInt[i] == int {
			return i
		}
	}
	return -1
}
