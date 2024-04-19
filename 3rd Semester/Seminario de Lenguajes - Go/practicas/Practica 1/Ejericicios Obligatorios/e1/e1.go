/*
Realice un programa que reciba una frase e imprima en pantalla la
misma frase reemplazando las ocurrencias de “jueves” por “martes”
respetando las letras minúsculas o mayúsculas de la palabra original en
su posición correspondiente. Por ejemplo, se reemplaza “Jueves” por
“Martes” o “jueveS” por “marteS”
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

// funcion encargada de recibir un char, identificar su capitalizacion, y devolverlo cambiado respestando capitalizacion
func toMartesB(b byte, l byte) byte {
	if unicode.IsUpper(rune(b)) {
		b = byte(unicode.ToUpper(rune(l)))
		return b
	} else {
		return l
	}
}

// funcion que separa la palabra jueves en un array de char, luego por el largo de la palabra entra a tomartesB
func toMartesS(s string) string {
	b := []byte(s)
	var mar string = "martes"
	for i := 0; i < 6; i++ {
		b[i] = toMartesB(b[i], mar[i])
	}
	//se hace un string del array de chars y se devuelve
	return string(b)
}

func main() {
	fmt.Print("Escriba una frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		phrase = scanner.Text()
	}
	//si contiene la palabra jueves, sigue el programa, sino termina
	if strings.Contains(strings.ToLower(phrase), "jueves") {
		//creamos un array compuesto de cada palabra separada por un " "
		words := strings.Fields(phrase)
		//vamos palabra por palabra identificando si es en realidad la palabra jueves, sin importar sus mayusculas.
		for i := 0; i < len(words); i++ {
			if strings.ToLower(words[i]) == "jueves" {
				//si lo es, transforma esa palabra del array en martes con la misma capitalizacion
				words[i] = toMartesS(words[i])
			}
		}
		//joinea el array en un string para imprimirlo
		fmt.Println("La nueva frase es: ", strings.Join(words, " "))
	} else {
		fmt.Println("No contiene la palabra \"jueves\".")
	}

}
