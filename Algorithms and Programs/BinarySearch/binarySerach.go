// Binary search

package main

import "fmt"

func main() {

	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	index := BinarySearch(sortedArray, target)

	if index != -1 {
		fmt.Printf("Objetivo %d encontrado en el indice %d\n", target, index)
	} else {
		fmt.Printf("Objetivo %d no encontrado en el indice\n", target)
	}
}

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
