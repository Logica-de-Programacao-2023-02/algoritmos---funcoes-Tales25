//Escreva uma função que calcule a média de uma lista de números e retorne um erro caso
//a lista esteja vazia.

package main

import "fmt"

func calculateMedia(list []float64) (float64, error) {
	var sum, i, media float64

	for i = 0; int(i) < len(list); i++ {
		sum += list[int(i)]
	}

	if len(list) == 0 {
		return 0, fmt.Errorf("list cannot be nil")
	}

	media = sum / i
	return media, nil
}

func main() {
	list := []float64{1, 1, 2, 3, 5, 8, 13, 21, 34}
	media, err := calculateMedia(list)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("A média da lista é: %.2f", media)
	}
}
