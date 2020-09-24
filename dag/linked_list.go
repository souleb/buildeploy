package dag

type LinkNode struct {
	next  *LinkNode
	value interface{}
}

type LinkedList struct {
	root *LinkNode
}

func NewLinkedList(value interface{}) *LinkedList {
	return &LinkedList{
		root: &LinkNode{
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
		current.next = &LinkNode{
			value: value,
		}

	} else {
		l.root = &LinkNode{
			value: value,
		}
	}
}

func (l *LinkedList) Add(index int, value interface{}) {
	node := &LinkNode{
		value: value,
	}
	currentIndex := 0
	var previous *LinkNode
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

func (l *LinkedList) Remove() interface{} {
	previousHead := l.root
	l.root = l.root.next
	return previousHead.value
}

func (l *LinkedList) Reverse() {
	current := l.root
	var previous, temp *LinkNode
	for current.next != nil {
		temp = current.next
		current.next = previous
		current = temp
	}
	current.next = previous
	l.root = current
}
