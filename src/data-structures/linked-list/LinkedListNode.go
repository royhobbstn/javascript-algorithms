package linkedlist

// Node a node in a LinkedList
type Node struct {
	Value interface{}
	Next  *Node
}

// NewNode create a new Node
func NewNode(value interface{}, next *Node) *Node {
	return &Node{Value: value, Next: next}
}
