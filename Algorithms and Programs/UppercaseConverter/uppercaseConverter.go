// Uppercase converter

package main

import (
	"fmt"
	"strings"
)

func main() {

	var text string

	fmt.Printf("Ingrese el texto que desea convertir a mayusculas: ")
	fmt.Scanln(&text)

	textUpperCase := upperCaseConverter(text)

	fmt.Printf("Texto en mayusculas: %s", textUpperCase)

}

func upperCaseConverter(text string) string {

	text = strings.ToUpper(text)
	textUpperCase := text

	return textUpperCase

}
