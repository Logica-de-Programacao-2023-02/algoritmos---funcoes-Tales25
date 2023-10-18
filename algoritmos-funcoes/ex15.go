//Crie uma função que receba um número inteiro e um slice de inteiros como parâmetros e retorne verdadeiro se o número
//estiver presente no slice e falso caso contrário. Caso o slice esteja vazio, retorne um erro.

package main

import "fmt"

func main() {
	var num int
	sliceInt := []int{1, 2, 4, 8, 16, 32, 64, 128}

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	fmt.Println(sliceInt)

	numberFound, err := FindNumberInSlice(num, sliceInt)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(numberFound)
	}
}

func FindNumberInSlice(num int, sliceInt []int) (bool, error) {
	if len(sliceInt) == 0 {
		return false, fmt.Errorf("slice cannnot be empty")
	}

	for _, x := range sliceInt {
		if num == x {
			return true, nil
		}
	}

	return false, nil
}
