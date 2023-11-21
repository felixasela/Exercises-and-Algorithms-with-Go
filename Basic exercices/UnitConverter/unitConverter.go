// Write a program that converts centimeters to inches. One inch is equal to 2.54 centimeters.

package main

import (
	"fmt"
)

func main() {

	var centimeter float64

	fmt.Printf("Ingrese la longitud en centimetros que desea convertir: ")
	fmt.Scanln(&centimeter)

	totalConvert := converter(centimeter)

	fmt.Printf("%.2fcm son equivalentes a %.2fin", centimeter, totalConvert)

}

func converter(centimeter float64) float64 {
	const inch = 2.54
	totalConvert := centimeter * inch
	return totalConvert
}
