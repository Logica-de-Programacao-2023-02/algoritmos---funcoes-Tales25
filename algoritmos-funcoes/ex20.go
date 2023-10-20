//Escreva uma função que receba um slice de strings como parâmetro e retorne um novo slice contendo
//apenas as strings que possuem mais de 5 caracteres. Caso o slice esteja vazio, retorne um erro.

package main

import "fmt"

func main() {
	sliceStr := []string{"Tales", "Tata", "Eduardo", "Duds", "Breno"}

	newSliceStr, err := WordsWithLen5More(sliceStr)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(newSliceStr)
	}
}

func WordsWithLen5More(sliceStr []string) ([]string, error) {
	var newSliceStr []string

	if len(sliceStr) == 0 {
		return nil, fmt.Errorf("slice cannot be empty")
	}

	for _, word := range sliceStr {
		if len(word) >= 5 {
			newSliceStr = append(newSliceStr, word)
		}
	}

	return newSliceStr, nil
}
