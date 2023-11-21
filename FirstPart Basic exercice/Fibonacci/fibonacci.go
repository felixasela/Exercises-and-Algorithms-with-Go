// Write a program that generates the Fibonacci sequence.

package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Printf("Ingrese la cantidad de términos de la secuencia de fibonacci: ")
	fmt.Scanln(&n)

	fibSequence := fibonacci(n)

	fmt.Printf("Secuencia de Fibonacci con %d términos %d", n, fibSequence)

}

func fibonacci(n int) []int {

	fibSequence := []int{0, 1}

	for len(fibSequence) < n {
		nextTerm := fibSequence[len(fibSequence)-1] + fibSequence[len(fibSequence)-2]
		fibSequence = append(fibSequence, nextTerm)
	}

	return fibSequence

}
