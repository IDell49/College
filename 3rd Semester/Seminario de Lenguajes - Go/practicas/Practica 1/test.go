package main

import "fmt"

func main() {
	a := [6]string{"a", "b", "c", "d", "e", "f"}
	s := a[2:4]
	fmt.Println(a, "\n", s)
	fmt.Println(cap(s))
}
