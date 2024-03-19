package main

import "fmt"

var z int = 1
var q int = 0
var sum int = 0

const max int = 250

func main() {
	for q != 2 {
		if z%2 == 0 {
			sum = sum + z
			q++
		}
		z++
	}
	fmt.Println("the sum is: ", sum)
}
