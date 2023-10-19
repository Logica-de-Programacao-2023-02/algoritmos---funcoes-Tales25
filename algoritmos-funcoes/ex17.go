//Crie uma função que receba um slice de strings como parâmetro e retorne uma nova string
//com as strings em ordem alfabética. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"fmt"
	"strings"
)

func main() {
	sliceStr := []string{"Eduardo", "Elsa", "Israel", "Tales", "Sidia", "Breno", "Lia", "Giulia"}

	strAlphabeticalOrder, err := AlphabeticalOrder(sliceStr)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(strAlphabeticalOrder)
	}
}

func AlphabeticalOrder(sliceStr []string) (string, error) {
	var str string

	if len(sliceStr) == 0 {
		return "", fmt.Errorf("slice cannot be empty")
	}

	for i := 0; i < len(sliceStr); i++ {
		for j := 0; j < len(sliceStr); j++ {
			sliceStr[i] = strings.ToLower(sliceStr[i])
			sliceStr[j] = strings.ToLower(sliceStr[j])

			if sliceStr[i][0] < sliceStr[j][0] {
				sliceStr[i], sliceStr[j] = sliceStr[j], sliceStr[i]
			} else if sliceStr[i][0] == sliceStr[j][0] && sliceStr[i][1] < sliceStr[j][1] {
				sliceStr[i], sliceStr[j] = sliceStr[j], sliceStr[i]
			}
		}
	}

	str = strings.Join(sliceStr, " ")

	return str, nil
}
