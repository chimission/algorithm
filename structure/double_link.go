// 双向链表

package structure

import "fmt"

type DoubleLinkNode struct {
	data int
	next *DoubleLinkNode
	prev *DoubleLinkNode
}

type DoubleLink struct {
	head *DoubleLinkNode
	tail *DoubleLinkNode
	len  int
}

func InitDoubleLink(data int) *DoubleLink {
	node := &DoubleLinkNode{data: data, next: nil, prev: nil}
	return &DoubleLink{head: node, tail: node, len: 1}
}

func (doubleLink *DoubleLink) AddFrontNode(data int) {
	newNode := &DoubleLinkNode{data: data}
	if doubleLink.head == nil {
		doubleLink.head = newNode
		doubleLink.tail = newNode
	} else {
		newNode.next = doubleLink.head
		doubleLink.head.prev = newNode
		doubleLink.head = newNode
	}

	doubleLink.len++
}

func (doubleLink *DoubleLink) AddEndNode(data int) {
	newNode := &DoubleLinkNode{data: data}
	if doubleLink.head == nil {
		doubleLink.head = newNode
		doubleLink.tail = newNode
	} else {
		currentNode := doubleLink.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
		newNode.prev = currentNode
	}
	doubleLink.len++
}

func (doubleLink *DoubleLink) Insert(data, index int) {
	if doubleLink.len < index-1 {
		return
	}
	newNode := &DoubleLinkNode{data: data}
	currentIndex := 0
	currentNode := doubleLink.head
	for currentNode.next != nil {
		currentNode = currentNode.next
		currentIndex++
		if currentIndex == index-1 {
			currentNode.next.prev = newNode
			newNode.next = currentNode.next
			newNode.prev = currentNode
			currentNode.next = newNode
			doubleLink.len++
			return
		}
	}

}

func (doubleLinkNode *DoubleLinkNode) Print() {
	fmt.Printf("=%d=", doubleLinkNode.data)
}

func (doubleLink *DoubleLink) Print() {
	fmt.Print("header=>")
	currentNode := doubleLink.head
	if currentNode != nil {
		currentNode.Print()
		for currentNode.next != nil {
			currentNode = currentNode.next
			currentNode.Print()
		}
	}
	fmt.Println("<=tail")
}
