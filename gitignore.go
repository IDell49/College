package main

import (
	"fmt"
	"os"
	"strings"
)

var ignore string

func ignoreFiller()

func spaceFiller(s []string) string {
	s := 
}

func main() {
	ignore, _ := os.ReadFile("./.gitignore")
	ignoreFiller(string(ignore))
	spaceFiller(strings.Fields(string(ignore)))
	os.WriteFile("./gitignore",ignore)
	fmt.Println("All done!")
}
