package main

import (
	"fmt"
)

var x int

func main() {
	fmt.Scanln(&x)
	if x < -18 {
		x = -x
	} else if x <= -1 {
		x = x % 4
	} else if x < 20 {
		x = x * x
	} else {
		x = x * -1.0
	}
	fmt.Println("the result is: ", x)
}
