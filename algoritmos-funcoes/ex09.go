//Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas na string.
//Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro.

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

	sliceWords, err := WordsSlice(str)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(sliceWords)
	}
}

func WordsSlice(str string) ([]string, error) {
	sliceWords := strings.Split(str, " ")

	if len(str) == 0 {
		return nil, fmt.Errorf("empty string/slice")
	}

	return sliceWords, nil
}
