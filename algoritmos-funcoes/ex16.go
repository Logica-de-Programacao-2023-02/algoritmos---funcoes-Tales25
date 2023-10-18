//Escreva uma função que receba uma string como parâmetro e retorne uma nova string com
//todas as vogais substituídas por "*". Caso a string seja vazia, retorne um erro.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	str = scanner.Text()

	newStr, err := ReplaceVowels(str)

	if err == nil {
		fmt.Print(newStr)
	} else {
		fmt.Print(err)
	}
}

func ReplaceVowels(str string) (string, error) {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, "a", "*")
	str = strings.ReplaceAll(str, "e", "*")
	str = strings.ReplaceAll(str, "i", "*")
	str = strings.ReplaceAll(str, "o", "*")
	str = strings.ReplaceAll(str, "u", "*")

	if len(str) == 0 {
		return "", fmt.Errorf("string cannot be empty")
	}

	return str, nil
}
