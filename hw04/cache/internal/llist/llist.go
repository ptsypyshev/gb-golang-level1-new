package llist

import "time"

type LinkedList interface {
	GetHead() *Node
	PopHead() *Node
	Append(n *Node)
}

type Node struct {
	updated time.Time
	key     string
	next    *Node
}

func (n *Node) Updated() time.Time {
	return n.updated
}

func (n *Node) Key() string {
	return n.key
}

func NewNode(u time.Time, k string) *Node {
	return &Node{
		updated: u,
		key:     k,
	}
}

type LinkedListImpl struct {
	head *Node
	tail *Node
}

func (l *LinkedListImpl) GetHead() *Node {
	return l.head
}

func (l *LinkedListImpl) PopHead() *Node {
	cur := l.head
	l.head = l.head.next
	return cur
}

func (l *LinkedListImpl) Append(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
	}

	l.tail.next = n
	l.tail = n
}
