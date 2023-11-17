package main

import (
	"fmt"
	"math"
)

func circleOperation(radius float64) (float64, float64) {
	areaCircle := math.Pi * math.Pow(radius, 2)
	perimeterCircle := 2 * radius * math.Pi
	return areaCircle, perimeterCircle
}

func main() {

	var radius float64

	fmt.Println("Ingrese el radio del círculo: ")
	fmt.Scanln(&radius)

	areaCircle, perimeterCircle := circleOperation(radius)

	fmt.Println("Área del círculo: ", areaCircle)
	fmt.Println("Perímetro del círculo: ", perimeterCircle)
}
