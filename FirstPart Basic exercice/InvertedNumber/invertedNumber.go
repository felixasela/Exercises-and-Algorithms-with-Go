/* Write a program that prompts the user for a three-digit integer and outputs the digits in reverse order */

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var number int

	fmt.Printf("Ingrese un numero de 3 digitos: ")
	fmt.Scanln(&number)

	invertNumber := invert(number)

	fmt.Printf("El numero invertido es: %d", invertNumber)

}

func invert(number int) int {
	numberStr := strconv.Itoa(number)
	character := []rune(numberStr)

	for i, j := 0, len(character)-1; i < j; i, j = i+1, j-1 {
		character[i], character[j] = character[j], character[i]
	}

	invertNumberStr := string(character)
	invertNumber, _ := strconv.Atoi(invertNumberStr)

	return invertNumber
}
