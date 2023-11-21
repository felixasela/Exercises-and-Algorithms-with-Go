// Write a program that outputs the decimal part of a real number input by the user.

package main

import (
	"fmt"
	"math"
)

func main() {

	var number float64

	fmt.Printf("Ingrese un numero decimal: ")
	fmt.Scanln(&number)

	decimalNumber := decimalPart(number)

	fmt.Printf("La parte decimal es: %f", decimalNumber)

}

func decimalPart(number float64) float64 {

	decimalNumber := number - math.Floor(number)
	return decimalNumber
}
