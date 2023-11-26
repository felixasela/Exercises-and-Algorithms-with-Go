// is it a full binary tree or not

package main

import (
	"fmt"
)

func main() {

	fullTree := &TreeNode{Val: 1}
	fullTree.Left = &TreeNode{Val: 2}
	fullTree.Right = &TreeNode{Val: 3}
	fullTree.Left.Left = &TreeNode{Val: 4}
	fullTree.Left.Right = &TreeNode{Val: 5}
	fullTree.Right.Left = &TreeNode{Val: 6}
	fullTree.Right.Right = &TreeNode{Val: 7}

	nonFullTree := &TreeNode{Val: 1}
	nonFullTree.Left = &TreeNode{Val: 2}
	nonFullTree.Right = &TreeNode{Val: 3}
	nonFullTree.Left.Left = &TreeNode{Val: 4}
	nonFullTree.Left.Right = &TreeNode{Val: 5}
	nonFullTree.Right.Left = &TreeNode{Val: 6}

	fmt.Println("Es el primer árbol un árbol binario completo?", isFullBinaryTree(fullTree))
	fmt.Println("Es el segundo árbol un árbol binario completo?", isFullBinaryTree(nonFullTree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isFullBinaryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if (root.Left == nil && root.Right == nil) || (root.Left != nil && root.Right != nil) {
		return isFullBinaryTree(root.Left) && isFullBinaryTree(root.Right)
	}

	return false
}
