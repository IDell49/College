package main

import "fmt"

func main() {
	var x, y, r float32
	fmt.Scanf("%d %d", &x, &y)
	if x > y {
		r = x / y
	} else {
		r = y / x
	}
	fmt.Print("el resul es: ", r, ". los numeros eran x: ", x, " y ", y)
}
