//Crie uma função que receba um número inteiro como parâmetro e retorne a soma de todos os seus dígitos.
//Caso o número seja negativo, retorne um erro.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		sumNumbers int
		num        int
		err        error
	)

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	sumNumbers, err = SumDigits(num)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(sumNumbers)
	}
}

func SumDigits(x int) (int, error) {
	var sumNumbers, numToInt int

	numToStr := strconv.Itoa(x)

	if x < 0 {
		return 0, fmt.Errorf("cannot be less than 0")
	}

	for i := 0; i < len(numToStr); i++ {
		//de string para inteiro pode ocorrer um erro caso tenha algum caractere diferente de algums numero então...
		//precisa de 2 variaveis para receber, normalmente o valor que será inteiro e um error (mas não nesse, pq...
		numToInt, _ = strconv.Atoi(string(numToStr[i])) //já é garantido que seja int)

		sumNumbers += numToInt
	}

	return sumNumbers, nil
}
