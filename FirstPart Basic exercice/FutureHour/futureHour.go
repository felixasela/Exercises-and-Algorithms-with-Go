/* Write a program that asks the user for the current time on the clock and an integer number of hours
that indicates the future time on the clock after the specified time */

package main

import (
	"fmt"
)

func main() {

	var (
		currentHour int8
		inputHour   int8
	)

	fmt.Printf("Ingrese la hora actual: ")
	fmt.Scanln(&currentHour)

	fmt.Printf("Ingrese las horas que desea agregar a la hora actual: ")
	fmt.Scanln(&inputHour)

	finalHour := sumHour(currentHour, inputHour)

	fmt.Printf("La hora proxima en formato militar son las: %d horas", finalHour)

}

func sumHour(currentHour int8, inputHour int8) int8 {

	finalHour := (currentHour + inputHour) % 24
	return finalHour
}
