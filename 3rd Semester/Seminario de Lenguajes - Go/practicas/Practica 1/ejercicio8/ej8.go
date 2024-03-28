/*
8. Realizar un programa que lea el punto cardinal (como caracter o
string) del cual viene el viento (‘N’, ‘S’, ‘E’, ‘O’) y envíe a la salida
estándar hacia cuál se dirigiría.
Sub-objetivo: Uso de case con la opción por default. E/S
caracteres o strings.
a. ¿Cómo se escribe el default en el case de otros lenguajes?
*/

package main

import "fmt"

func main() {
	fmt.Println("escriba desde donde viene el viento: ")
	var punto string
	fmt.Scanln(&punto)
	switch {
	case punto == "N":
		fmt.Println("el viento se dirige hacia el Sur")

	case punto == "S":
		fmt.Println("el viento se dirige hacia el Norte")

	case punto == "E":
		fmt.Println("el viento se dirige hacia el Oeste")

	default:
		fmt.Println("el viento se dirige hacia el Este")
	}
}
