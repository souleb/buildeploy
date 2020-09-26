package dag

import "fmt"

type linkNode struct {
	next  *linkNode
	value interface{}
}

type LinkedList struct {
	root *linkNode
}

func NewLinkedList(value interface{}) *LinkedList {
	return &LinkedList{
		root: &linkNode{
			value: value,
		},
	}
}

func (l *LinkedList) Append(value interface{}) {
	if l.root != nil {
		current := l.root
		for current.next != nil {
			current = current.next
		}
		current.next = &linkNode{
			value: value,
		}

	} else {
		l.root = &linkNode{
			value: value,
		}
	}
}

func (l *LinkedList) Add(index int, value interface{}) {
	node := &linkNode{
		value: value,
	}
	currentIndex := 0
	var previous *linkNode
	current := l.root
	for currentIndex < index {
		previous = current
		current = current.next
		currentIndex++
	}
	node.next = current
	previous.next = node

}

func (l *LinkedList) Length() int {
	length := 0
	current := l.root
	for current != nil {
		length++
		current = current.next
	}

	return length
}

func (l *LinkedList) Remove(value interface{}) error {
	var previous *linkNode
	current := l.root
	for current.next != nil {
		if current == value {
			previous.next = current.next
			return nil
		}
		previous = current
		current = current.next
	}

	return fmt.Errorf("Value not found")
}

func (l *LinkedList) Reverse() {
	current := l.root
	var previous, temp *linkNode
	for current.next != nil {
		temp = current.next
		current.next = previous
		current = temp
	}
	current.next = previous
	l.root = current
}
