/*
Las temperaturas de los pacientes de un hospital se dividen en 3
grupos: alta (mayor de 37.5), normal (entre 36 y 37.5) y baja
(menor de 36). Se deben leer 10 temperaturas de pacientes e
informar el porcentaje de pacientes de cada grupo. Luego se
debe imprimir el promedio entero entre la temperatura máxima y
la temperatura mínima.
a. ¿Se puede utilizar el case para tipos reales en otros
lenguajes?
b. ¿ Cómo se realizan las conversiones entre reales (punto
flotante) y enteros en otros lenguajes ?
Sub-objetivo: El tipado fuerte, usar casting. Operaciones y
E/S con float. Casting en otros lenguajes
*/
package main

import "fmt"

const high float32 = 37.5
const low float32 = 36

var degrees float32
var degreesavg float32 = 0
var highq float32 = 0
var medq float32 = 0
var lowq float32 = 0

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("paso:", i)
		fmt.Scanln(&degrees)
		switch {
		case degrees > high:
			highq++
			degreesavg = degreesavg + degrees
		case degrees < low:
			lowq++
			degreesavg += degrees
		default:
			medq++
		}
		fmt.Println("cantidades: ", highq, " ", medq, " ", lowq)
	}
	fmt.Println("el porcentaje por cada grupo es:\nalto: ", highq/10*100, "%\nmedio: ", medq/10*100, "%\nbajo:", lowq/10*100, "%")
	fmt.Println("el promedio entre la tempe alta y baja es: ", int(degreesavg/(highq+lowq)))
}
