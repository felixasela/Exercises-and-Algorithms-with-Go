// Write a program that calculates the average of four notes input by the user.

package main

import (
	"fmt"
)

func main() {
	var (
		noteOne   float64
		noteTwo   float64
		noteThree float64
		notefour  float64
	)

	fmt.Printf("Ingrese la nota 1: ")
	fmt.Scanln(&noteOne)

	fmt.Printf("Ingrese la nota 2: ")
	fmt.Scanln(&noteTwo)

	fmt.Printf("Ingrese la nota 3: ")
	fmt.Scanln(&noteThree)

	fmt.Printf("Ingrese la nota 4: ")
	fmt.Scanln(&notefour)

	average := calculateAverage(noteOne, noteTwo, noteThree, notefour)

	fmt.Printf("El promedio de notas es %.2f", average)

}

func calculateAverage(noteOne float64, noteTwo float64, noteThree float64, notefour float64) float64 {
	average := (noteOne + noteTwo + noteThree + notefour) / 4
	return average
}
