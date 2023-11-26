// Maximum Path Sum in a Binary Tree

package main

import (
	"fmt"
	"math"
)

func main() {

	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 10}
	root.Left.Left = &TreeNode{Val: 20}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: -25}
	root.Right.Right.Left = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 4}

	result := maxPathSum(root)
	fmt.Println("Suma máxima de rutas en el árbol binario:", result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSumHelper(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	leftSum := maxPathSumHelper(root.Left, result)
	rightSum := maxPathSumHelper(root.Right, result)

	currentSum := root.Val + leftSum + rightSum

	*result = int(math.Max(float64(*result), float64(currentSum)))

	return int(math.Max(0, float64(root.Val)+math.Max(float64(leftSum), float64(rightSum))))
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt64
	maxPathSumHelper(root, &result)
	return result
}
