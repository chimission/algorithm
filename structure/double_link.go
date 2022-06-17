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

func InitDoubleLink() *DoubleLink {
	return &DoubleLink{head: nil, tail: nil, len: 0}
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
	if doubleLink.tail == nil {
		doubleLink.head = newNode
		doubleLink.tail = newNode
	} else {
		newNode.prev = doubleLink.tail
		doubleLink.tail.next = newNode
		doubleLink.tail = newNode
	}
	doubleLink.len++
}

func (doubleLink *DoubleLink) InsertByHead(newNode *DoubleLinkNode, index int) {
	currentIndex := 0
	currentNode := doubleLink.head
	for currentNode.next != nil {
		if currentIndex == index-1 {
			currentNode.next.prev = newNode
			newNode.next = currentNode.next
			newNode.prev = currentNode
			currentNode.next = newNode
			doubleLink.len++
			return
		}
		currentNode = currentNode.next
		currentIndex++
	}

}

func (doubleLink *DoubleLink) InsertByTail(newNode *DoubleLinkNode, index int) {
	currentIndex := doubleLink.len
	currentNode := doubleLink.tail
	for currentNode.prev != nil {
		if currentIndex == index+1 {
			currentNode.prev.next = newNode
			newNode.prev = currentNode.prev
			newNode.next = currentNode
			currentNode.prev = newNode
			doubleLink.len++
			return
		}
		currentNode = currentNode.prev
		currentIndex--
	}

}

func (doubleLink *DoubleLink) Insert(data, index int) {
	if index < 0 || index > doubleLink.len {
		return
	}
	fmt.Println(doubleLink.len, index)
	if index == 0 {
		doubleLink.AddFrontNode(data)
	} else if index == doubleLink.len {
		doubleLink.AddEndNode(data)
	} else {
		newNode := &DoubleLinkNode{data: data}
		//  判断插入位置靠近head 还是 靠近tail 减少循环量
		m := doubleLink.len / 2
		if index <= m {
			doubleLink.InsertByHead(newNode, index)
		} else {
			doubleLink.InsertByTail(newNode, index)
		}
	}

}

func (doubleLinkNode *DoubleLinkNode) Print() {
	fmt.Printf("=%d=", doubleLinkNode.data)
}

func (doubleLink *DoubleLink) HeadPrint() {
	fmt.Print("header=")
	currentNode := doubleLink.head
	if currentNode != nil {
		currentNode.Print()
		for currentNode.next != nil {
			currentNode = currentNode.next
			currentNode.Print()
		}
	}
	fmt.Println("=tail")
}

func (doubleLink *DoubleLink) TailPrint() {
	fmt.Print("tail=")
	currentNode := doubleLink.tail
	if currentNode != nil {
		currentNode.Print()
		for currentNode.prev != nil {
			currentNode = currentNode.prev
			currentNode.Print()
		}
	}
	fmt.Println("=header")
}
