package main

import "fmt"

func main() {
	var (
		nom string
		ape string
		tel int
	)
	n, e := fmt.Scanf("%s %s %d", &nom, &ape, &tel)

	if e != nil {
		fmt.Printf("Error: %s", e)
	} else {
		fmt.Printf("Todo bien, recibimos %d argumentos: %s, %s, %d", n, nom, ape, tel)
	}
}
