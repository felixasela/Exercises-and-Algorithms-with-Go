// Delete a node from the linked list

package main

import "fmt"

func main() {
	linkedList := LinkedList{}

	linkedList.InsertAtEnd(1)
	linkedList.InsertAtEnd(2)
	linkedList.InsertAtEnd(3)

	fmt.Println("Lista enlazada antes de la eliminación:")
	linkedList.Display()

	linkedList.DeleteNode(2)

	fmt.Println("\nLista enlazada después de la eliminación:")
	linkedList.Display()
}

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) DeleteNode(data int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current.Next != nil && current.Next.Data != data {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}
