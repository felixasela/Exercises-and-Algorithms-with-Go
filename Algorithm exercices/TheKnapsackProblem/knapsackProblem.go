// The Knapsack problem

package main

import (
	"fmt"
)

func main() {
	items := []Item{
		{Weight: 2, Value: 10},
		{Weight: 5, Value: 30},
		{Weight: 8, Value: 40},
		{Weight: 2, Value: 5},
	}

	capacity := 10

	maxValue, selectedItems := knapsack(items, capacity)

	fmt.Printf("Valor maximo: %d\n", maxValue)
	fmt.Printf("Items sekeccionados: %v\n", selectedItems)
}

type Item struct {
	Weight int
	Value  int
}

func knapsack(items []Item, capacity int) (int, []Item) {
	n := len(items)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			item := items[i-1]

			if item.Weight <= w {
				include := item.Value + dp[i-1][w-item.Weight]
				exclude := dp[i-1][w]
				dp[i][w] = max(include, exclude)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	selectedItems := make([]Item, 0)
	i, j := n, capacity
	for i > 0 && j > 0 {
		if dp[i][j] != dp[i-1][j] {
			selectedItems = append(selectedItems, items[i-1])
			j -= items[i-1].Weight
		}
		i--
	}

	return dp[n][capacity], selectedItems
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
