package singlelinkedlist

type Node struct {
	Value    any
	Previous *Node
	Next     *Node
}

type SingleLinkedList struct {
	Head *Node
	Tail *Node
}

func (s *SingleLinkedList) Append(value any) {
	newNode := &Node{Value: value}

	if s.Tail == nil {
		s.Head = newNode
		s.Tail = newNode
	} else {
		newNode.Previous = s.Tail
		s.Tail.Next = newNode
		s.Tail = newNode
	}
}
