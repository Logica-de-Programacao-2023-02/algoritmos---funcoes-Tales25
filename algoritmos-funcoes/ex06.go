//Escreva uma função que receba um slice de strings como parâmetro e retorne uma string com todas as strings
//concatenadas e separadas por vírgulas. Caso o slice esteja vazio, retorne um erro.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		strSlice             []string
		str, concatenatedStr string
		strQuantity          int
	)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Escolha a quantidade de palavras: ")
	//usei Scanln para "consumir uma quebra de linha", com o Scan normal o programa lia a primeira palavra como " "
	fmt.Scanln(&strQuantity)

	for i := 0; i < strQuantity; i++ {
		fmt.Printf("Escolha a palavra de número %d: ", i)
		scanner.Scan()
		str = scanner.Text()

		strSlice = append(strSlice, str)
	}

	concatenatedStr = concatenateStrings(strSlice)
	fmt.Printf("Transformando o slice de strings em uma única string temos:\n%s", concatenatedStr)
}

func concatenateStrings(strSlice []string) string {
	var concatenated string
	concatenated = strings.Join(strSlice, ", ")
	return concatenated
}
