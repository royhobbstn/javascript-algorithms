package linkedlist

// LinkedList a linked list data structure
type LinkedList struct {
	Head *Node
	Tail *Node
}

func (p *LinkedList) prepend(value interface{}) *LinkedList {
	// Make new node to be a head.
	newNode := NewNode(value, p.Head)
	p.Head = newNode

	// If there is no tail yet let's make new node a tail.
	if p.Tail == nil {
		p.Tail = newNode
	}

	return p
}

func (p *LinkedList) append(value interface{}) *LinkedList {
	newNode := NewNode(value, nil)

	// If there is no head yet let's make new node a head.
	if p.Head == nil {
		p.Head = newNode
		p.Tail = newNode
		return p
	}

	// Attach new node to the end of linked list.
	p.Tail.Next = newNode
	p.Tail = newNode
	return p
}

// func (p *LinkedList) delete(value interface{}) *Node {

// }

// func (p *LinkedList) find(value interface{}, callback func(){}) *Node {

// }

func (p *LinkedList) deleteTail() *Node {
	deletedTail := p.Tail

	if p.Head == p.Tail {
		// There is only one node in linked list.
		p.Head = nil
		p.Tail = nil
		return deletedTail
	}

	// If there are many nodes in linked list...

	// Rewind to the last node and delete "next" link for the node before the last one.
	currentNode := p.Head
	for currentNode.Next != nil {
		if currentNode.Next.Next == nil {
			currentNode.Next = nil
		} else {
			currentNode = currentNode.Next
		}
	}

	p.Tail = currentNode

	return deletedTail
}

func (p *LinkedList) deleteHead() *Node {
	if p.Head == nil {
		return nil
	}

	deletedHead := p.Head

	if p.Head.Next != nil {
		p.Head = p.Head.Next
	} else {
		p.Head = nil
		p.Tail = nil
	}

	return deletedHead
}

// func (p *LinkedList) fromArray() *Node {

// }

// func (p *LinkedList) toArray() *Node {

// }

// func (p *LinkedList) toString() *Node {

// }

// func (p *LinkedList) reverse() *Node {

// }
