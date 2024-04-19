/*
Realice un programa que reciba una palabra como argumento y lee de
la entrada una frase. Luego, el programa debe imprimir la frase que
leyó con cada una de las ocurrencias de la palabra con las mayúsculas
y minúsculas invertidas. Por ejemplo, si la frase es:
“Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO”
y la palabra es “PEQUEÑO” entonces el programa imprimirá:
“Parece PEQueÑO, pero no es tan PEQUEñO el pequeño”
Tenga en cuenta que la palabra a buscar puede ser ingresada con
mayúsculas y minúsculas mezcladas.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var phrase, prevWord, targetWord string

// funcion que separa la palabra recibida en un array de runas, luego itera ese array y va reemplazando las letras una por una
func toAutomovilS(i string, t string) string {
	input := []rune(i)
	// Convertir la cadena de reemplazo en una secuencia de runes
	target := []rune(t)
	// Iterar sobre los caracteres de la entrada y la cadena de reemplazo
	for i := 0; i < len(input); i++ {
		// Cambiar la capitalización del caracter según la palabra original
		if unicode.IsUpper(input[i]) {
			input[i] = unicode.ToUpper(target[i])
		} else {
			input[i] = target[i]
		}
	}
	// Devolver la cadena resultante
	return string(input)
}

func getInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	word, _ := reader.ReadString('\n')

	return strings.TrimSpace(word)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	phrase := getInput("Escribi una frase: ", reader)
	prevWord := getInput("Palabra a cambiar: ", reader)
	targetWord := getInput("Por cual la deseás cambiar: ", reader)

	// Si la frase contiene la palabra original, y si esa palabra coincide en largo con el target sigue, sino termina.
	if (strings.Contains(strings.ToLower(phrase), prevWord)) && (len([]rune(prevWord)) == len([]rune(targetWord))) {
		words := strings.Fields(phrase)
		for i := 0; i < len(words); i++ {
			if strings.ToLower(words[i]) == prevWord {
				words[i] = toAutomovilS(words[i], targetWord)
			}
		}
		fmt.Println("The phrase ended up like:\"", strings.Join(words, " "), "\"")

	} else {
		if strings.Contains(strings.ToLower(phrase), prevWord) {
			fmt.Println("No contiene la palabra \"", prevWord, "\".")
		} else {
			fmt.Println("Por favor, ingrese dos palabras con mismo largo")
		}
	}

}
