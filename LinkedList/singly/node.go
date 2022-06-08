package llsingly

type Node[T Types] struct {
	Value T
	Next  *Node[T]
}

func InitNode[T Types](value T) *Node[T] {
	return &Node[T]{
		value,
		nil,
	}
}
