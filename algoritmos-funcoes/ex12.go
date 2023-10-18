//Escreva uma função que receba um slice de strings como parâmetro e retorne uma string contendo apenas
//as strings que começam com uma letra maiúscula. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"fmt"
	"strings"
)

func main() {
	sliceStr := []string{"Tales", "breno", "Israel", "sidia", "tales", "Breno", "israel", "Sidia"}

	strCapitalized, err := InitialCapitalized(sliceStr)

	if err == nil {
		fmt.Print(strCapitalized)
	} else {
		fmt.Print(err)
	}
}

func InitialCapitalized(sliceStr []string) (string, error) {
	var newStr string
	var sliceCapitalized []string

	if len(sliceStr) == 0 {
		return "", fmt.Errorf("empty slice")
	}

	for _, words := range sliceStr {
		if words[0] >= 'A' && words[0] <= 'Z' {
			sliceCapitalized = append(sliceCapitalized, words)
		}
	}

	newStr = strings.Join(sliceCapitalized, " ")

	return newStr, nil
}
