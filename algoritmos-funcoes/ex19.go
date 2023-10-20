//Crie uma função que receba um número inteiro como parâmetro e retorne um novo slice contendo todos os números
//primos menores ou iguais a ele. Caso o número seja negativo, retorne um erro.

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	primes, err := FindSmallerPrimes(num)

	if err == nil {
		fmt.Print(primes)
	} else {
		fmt.Print(err)
	}
}

func FindSmallerPrimes(num int) ([]int, error) {
	var slicePrimes []int
	var isPrime bool

	if num <= 1 {
		return nil, fmt.Errorf("num cannnot be one or less")
	}

	if num == 2 {
		slicePrimes = append(slicePrimes, 2)
	} else if num == 3 {
		slicePrimes = append(slicePrimes, 2, 3)
	} else if num == 5 {
		slicePrimes = append(slicePrimes, 2, 3, 5)
	} else if num >= 7 {
		slicePrimes = append(slicePrimes, 2, 3, 5, 7)
	}

	for i := 4; i < num; i++ {
		isPrime = true

		if i%2 == 0 || i%3 == 0 || i%5 == 0 || i%7 == 0 {
			isPrime = false
		}

		if isPrime {
			slicePrimes = append(slicePrimes, i)
		}
	}

	return slicePrimes, nil
}
