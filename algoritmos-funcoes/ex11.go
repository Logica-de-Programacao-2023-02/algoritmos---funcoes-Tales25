//Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro se ele for um número primo
//e falso caso contrário. Caso o número seja negativo, retorne um erro.

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	isPrime, err := IsPrime(num)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(isPrime)
	}
}

func IsPrime(num int) (bool, error) {
	if num <= 1 {
		return false, fmt.Errorf("cannot be one or less")
	}

	if num == 2 || num == 3 {
		return true, nil
	}

	if num%2 == 0 || num%3 == 0 {
		return false, nil
	}

	return true, nil
}
