package tests

import (
	llsingly "ds-go/LinkedList/singly"
	"testing"
)

func TestLinkedListAppend(t *testing.T) {
	ll := llsingly.LinkedList[int]{}
	ll.Append(5)
	if ll.Head.Value != 5 {
		t.Errorf("linked list: head value is wrong")
	}
	ll.Print()
}
