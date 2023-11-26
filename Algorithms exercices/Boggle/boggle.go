// Boggle game

package main

import (
	"fmt"
)

var dictionary = map[string]bool{
	"CAR": true,
	"CAT": true,
	"RAT": true,
	"BAT": true,
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

func insertWord(root *TrieNode, word string) {
	node := root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func boggleDFS(board [][]rune, visited [][]bool, row, col int, currentWord string, trieNode *TrieNode, result map[string]bool) {
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || visited[row][col] {
		return
	}

	char := board[row][col]
	currentWord += string(char)
	visited[row][col] = true

	if node, ok := trieNode.children[char]; ok {
		if node.isEnd {
			result[currentWord] = true
		}

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				boggleDFS(board, visited, row+i, col+j, currentWord, node, result)
			}
		}
	}

	visited[row][col] = false
}

func findWords(board [][]rune, trieRoot *TrieNode) map[string]bool {
	result := make(map[string]bool)
	visited := make([][]bool, len(board))

	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			boggleDFS(board, visited, i, j, "", trieRoot, result)
		}
	}

	return result
}

func main() {
	board := [][]rune{
		{'C', 'A', 'R'},
		{'E', 'T', 'P'},
		{'K', 'A', 'T'},
	}

	root := NewTrieNode()

	for word := range dictionary {
		insertWord(root, word)
	}

	foundWords := findWords(board, root)

	fmt.Println("Palabras encontradas:")
	for word := range foundWords {
		fmt.Println(word)
	}
}
