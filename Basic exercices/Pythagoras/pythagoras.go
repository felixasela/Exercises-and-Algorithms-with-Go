/* Write a program that prompts the user to input the lengths of the two cathetus, 'a' and 'b'
of a right-angled triangle, and outputs the length of the hypotenuse 'c' using the Pythagorean theorem */

package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		cathetusA float64
		cathetusB float64
	)

	fmt.Printf("Ingrese la longitud del cateto A: ")
	fmt.Scanln(&cathetusA)

	fmt.Printf("Ingrese la longitud del cateto B: ")
	fmt.Scanln(&cathetusB)

	cathetusC := pythagorasTheorem(cathetusA, cathetusB)

	fmt.Printf("La longitud del cateto C es de: %.2f", cathetusC)

}

func pythagorasTheorem(cathetusA float64, cathetusB float64) float64 {

	cathetusA = math.Pow(cathetusA, 2)
	cathetusB = math.Pow(cathetusB, 2)
	sumCathetus := cathetusA + cathetusB
	cathetusC := math.Sqrt(sumCathetus)

	return cathetusC
}
