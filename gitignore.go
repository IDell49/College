// El objetivo es que al ejecutarse:
// 1. Abra .gitignore
// 2. Sepa agregar todos los archivos NUEVOS de 100+mb sin repetir
// 3. Cambie todos los espacios con * para q sepa leerlo
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var ignore string

func ignoreFiller(new []string) []string {
	//exec.Command("find . -size +100M | cat >> .gitignore")
	cmd := exec.Command("find", ".", "-size", "+100M")
	output, _ := cmd.Output()
	newLines := strings.Split(string(output), "\n")

	// Iterate over each line
	for _, line := range newLines {
		// Check if the line starts with the prefix "./.git"
		if !strings.HasPrefix(line, "./.git") {
			// Keep the line if it does not start with the prefix
			line, _ = strings.CutPrefix(line, "./")
			new = append(new, line)
		}
	}
	return new
}

func spaceFiller(s []string) []string {
	//fill spaces with *
	for i := 0; i < len(s); i++ {
		//make a words array from each line of the document
		words := strings.Fields(s[i])
		//join the words with * instead of spaces
		filled := strings.Join(words, "*")
		s[i] = filled
	}
	return s
}

func main() {
	//read manual gitignores
	origin, _ := os.ReadFile("./manualignore.txt")
	manual := strings.Split(string(origin), "\n")
	//update auto-generated gitignore lines
	var newlines []string
	newlines = ignoreFiller(newlines)
	newlines = spaceFiller(newlines)
	//join both of them
	gitignore := strings.Join(append(manual, newlines...), "\n")
	//write into the file
	os.WriteFile(".gitignore", []byte(gitignore), 0644)
	fmt.Println("All done!")
}
