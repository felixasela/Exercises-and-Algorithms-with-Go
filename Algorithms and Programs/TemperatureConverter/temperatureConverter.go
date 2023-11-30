// Temperature converter

package main

import "fmt"

func main() {

	var celcius float64
	fmt.Printf("Ingrese la cantidad en grados celsius: ")
	fmt.Scanln(&celcius)

	temp := temperatureConverter(celcius)

	fmt.Printf("La temperatura en grados fahrenheit es de: %.2f", temp)

}

func temperatureConverter(temp float64) float64 {

	tempConverter := temp * 33.8

	return tempConverter

}
