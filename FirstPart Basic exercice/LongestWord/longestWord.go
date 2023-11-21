// Write a program that asks the user for two words and indicates which word is longer and by how many letters

package main

import (
	"fmt"
)

func main() {

	var word1, word2 string

	fmt.Printf("Ingrese la primera palabra: ")
	fmt.Scanln(&word1)

	fmt.Printf("Ingrese la segunda palabra: ")
	fmt.Scanln(&word2)

	longestWord(word1, word2)

}

func longestWord(word1 string, word2 string) {

	lengthWord1 := len(word1)
	lengthWord2 := len(word2)

	if lengthWord1 > lengthWord2 {
		fmt.Printf("La palabra 1 es mas larga que la 2 y tiene %d caracteres", lengthWord1)

	} else {
		fmt.Printf("La palabra 2 es mas larga que la 1 y tiene %d caracteres", lengthWord2)
	}

}
