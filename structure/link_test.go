package structure

import (
	"testing"
)

func TestNewLink(t *testing.T) {
	header := NewLinkNode(1)
	if header.Next != nil {
		t.Error("new link error")
	}
}

func TestLinkAdd(t *testing.T) {
	header := NewLinkNode(1)
	header.Add(2)
	header.Add(3)
	header.Add(4)
	header.Add(5)
	header.Add(6)
	if header.Len() != 6 {
		t.Errorf("LinkNode Add error, Get %d, Should %d", header.Len(), 6)
	}
}
