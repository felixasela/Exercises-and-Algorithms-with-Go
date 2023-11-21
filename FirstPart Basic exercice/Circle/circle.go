// Write a program that receives the radius of a circle as input and delivers its perimeter and area as output

package main

import (
	"fmt"
	"math"
)

func main() {

	var radius float64

	fmt.Printf("Ingrese el radio del círculo: ")
	fmt.Scanln(&radius)

	areaCircle, perimeterCircle := circleOperation(radius)

	fmt.Printf("Área del círculo: %.2f \n", areaCircle)
	fmt.Printf("Perímetro del círculo: %.2f \n", perimeterCircle)
}

func circleOperation(radius float64) (float64, float64) {
	areaCircle := math.Pi * math.Pow(radius, 2)
	perimeterCircle := 2 * radius * math.Pi
	return areaCircle, perimeterCircle
}
