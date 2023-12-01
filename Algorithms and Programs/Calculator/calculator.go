package main

import "fmt"

func main() {
	var (
		num1, num2, result float64
		operator           string
	)

	fmt.Printf("Ingrese el primer numero: ")
	fmt.Scanln(&num1)

	fmt.Printf("Ingrese el operador de la operacion que desea realizar (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Printf("Ingrese el segundo numero: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		result = addition(num1, num2)
		fmt.Printf("El resultado de la suma es: %.2f", result)
	case "-":
		result = subtraction(num1, num2)
		fmt.Printf("El resultado de la resta es: %.2f", result)
	case "*":
		result = multiplication(num1, num2)
		fmt.Printf("El resultado de la multiplicacion es: %.2f", result)
	case "/":
		result = division(num1, num2)
		fmt.Printf("El resultado de la division es: %.2f", result)
	}
}

func addition(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func subtraction(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func multiplication(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func division(num1 float64, num2 float64) float64 {

	if num1 != 0 && num2 != 0 {
		division := num1 / num2
		return division
	}

	fmt.Printf("No se puede dividir por 0 ")
	return 0
}
