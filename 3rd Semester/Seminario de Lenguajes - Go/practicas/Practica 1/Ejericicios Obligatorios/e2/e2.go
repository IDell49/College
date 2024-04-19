/*
Realice las modificaciones necesarias al ejercicio anterior para que en
lugar de reemplazar la palabra “jueves” por “martes” ahora se
reemplace “miércoles” por “automóvil”. Piense qué impacto tuvieron
esas modificaciones en el programa que había realizado.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var phrase string

// funcion que separa la palabra “miércoles” en un array de runas, luego itera ese array y va reemplazando las letras una por una
func toAutomovilS(i string) string {
	input := []rune(i)
	t := "automóvil"
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

func main() {
	//se escribe la frase de entrada
	fmt.Println("Escriba una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		phrase = scanner.Text()
	}

	//si contiene la palabra “miércoles”, sigue el programa, sino termina
	if strings.Contains(strings.ToLower(phrase), "miércoles") {

		//creamos un array compuesto de cada palabra separada por un " "
		words := strings.Fields(phrase)
		//vamos palabra por palabra identificando si es en realidad la palabra “miércoles”, sin importar sus mayusculas.
		for i := 0; i < len(words); i++ {
			if strings.ToLower(words[i]) == "miércoles" {
				//si lo es, transforma esa palabra del array en automóvil con la misma capitalizacion
				words[i] = toAutomovilS(words[i])
			}
		}
		//joinea el array en un string para imprimirlo
		fmt.Println("La nueva frase es: ", strings.Join(words, " "))
	} else {
		fmt.Println("No contiene la palabra \"miércoles\".")
	}

}
