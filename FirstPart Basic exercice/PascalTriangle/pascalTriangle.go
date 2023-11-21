// Write a program that generetes the pascal triangle

package main

import (
	"fmt"
)

func main() {
	var numRows int

	fmt.Print("Ingrese el numero de filas para el triangulo de pascal: ")
	fmt.Scanln(&numRows)

	pascalsTriangle := generatePascalsTriangle(numRows)

	fmt.Printf("Triangulo de pascal con %d filas:\n", numRows)
	printPascalsTriangle(pascalsTriangle)
}

func generatePascalsTriangle(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func printPascalsTriangle(triangle [][]int) {
	for _, row := range triangle {
		fmt.Println(row)
	}
}
