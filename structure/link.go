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

func (linkNode *LinkNode) Delete(index int) *LinkNode {
	if index == 0 {
		t := linkNode.Next
		linkNode.Next = nil
		return t
	}
	header := linkNode
	pre := linkNode
	for i := 0; pre != nil; i++ {
		if i == index-1 {
			t := pre.Next
			pre.Next = pre.Next.Next
			t.Next = nil
		}
		pre = pre.Next
	}
	return header
}

func (linkNode *LinkNode) Print() {
	for linkNode != nil {
		fmt.Print(linkNode.Data, "-->")
		linkNode = linkNode.Next
	}
}
