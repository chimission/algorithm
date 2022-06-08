// 单向链表

package structure

import "fmt"

type LinkNode struct {
	Data int64
	Next *LinkNode
}

func NewLinkNode(data int64) *LinkNode {
	header := new(LinkNode)
	header.Data = data
	return header
}

func (linkNode *LinkNode) Add(data int64) {
	for linkNode.Next != nil {
		linkNode = linkNode.Next
	}
	newLinkNode := NewLinkNode(data)
	linkNode.Next = newLinkNode

}

func (linkNode *LinkNode) Len() (lenght int) {

	for linkNode != nil {
		linkNode = linkNode.Next
		lenght++
	}
	return lenght
}

func (linkNode *LinkNode) Insert(data int64, index int) *LinkNode { //TODO judge length

	if index == 0 {
		newLinkNode := NewLinkNode(data)
		newLinkNode.Next = linkNode
		return newLinkNode
	} else {
		header := linkNode
		for i := 0; linkNode != nil; i++ {
			if i == index-1 {
				newLinkNode := NewLinkNode(data)
				t := newLinkNode.Next
				linkNode.Next = newLinkNode
				newLinkNode.Next = t
			}
			linkNode = linkNode.Next
		}
		return header
	}

}

func (linkNode *LinkNode) Delete(index int) *LinkNode {
	if index == 0 {
		pre := linkNode
		linkNode = linkNode.Next
		pre.Next = nil

	} else {
		pre := linkNode
		for i := 0; pre != nil; i++ {
			if i == index-1 {
				t := pre.Next
				pre.Next = pre.Next.Next
				t.Next = nil
			}
			pre = pre.Next
		}
	}

	return linkNode
}

func (linkNode *LinkNode) Print() {
	for linkNode != nil {
		fmt.Print(linkNode.Data, "-->")
		linkNode = linkNode.Next
	}
	fmt.Println("nil")
}
