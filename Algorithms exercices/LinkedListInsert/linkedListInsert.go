// Insert a Node in the Linked List

package main

import "fmt"

func main() {

	linkedList := LinkedList{}

	linkedList.InsertAtEnd(1)
	linkedList.InsertAtEnd(2)
	linkedList.InsertAtEnd(3)

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
