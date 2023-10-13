//Crie uma função que receba uma string e retorne a quantidade de vogais presentes.

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

	vowelsQuantity, err := calculateVowels(str)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("A frase tem %d vogais", vowelsQuantity)
	}
}

func calculateVowels(str string) (int, error) {
	var vowelsQuantity int
	var haveVowels bool
	haveVowels = false

	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "a", " ")
	str = strings.ReplaceAll(str, "e", " ")
	str = strings.ReplaceAll(str, "i", " ")
	str = strings.ReplaceAll(str, "o", " ")
	str = strings.ReplaceAll(str, "u", " ")

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			vowelsQuantity++
			haveVowels = true
		}
	}

	if !haveVowels {
		return 0, fmt.Errorf("there are no vowels")
	}

	return vowelsQuantity, nil
}
