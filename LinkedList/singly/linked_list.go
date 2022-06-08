package llsingly

import (
	"fmt"
)

type LinkedList[T Types] struct {
	Head *Node[T]
}

func InitLinkedList[T Types]() *LinkedList[T] {
	return &LinkedList[T]{
		nil,
	}
}

func (ll *LinkedList[T]) Append(values ...T) {
	if len(values) == 0 {
		return
	}

	if ll.Head == nil {
		ll.Head = InitNode(values[0])
	}

	node := ll.Head

	// iterates until find an empty next
	for node.Next != nil {
		node = node.Next
	}

	// Adds the values
	for _, val := range values[1:] {
		node.Next = InitNode(val)
		node = node.Next
	}
}

func (ll *LinkedList[T]) Print() error {
	node := ll.Head
	str := "["
	for node != nil {
		str += fmt.Sprint(node.Value)
		node = node.Next
		if node != nil {
			str += " "
		}
	}

	str += "]"
	fmt.Println(str)

	return nil
}
