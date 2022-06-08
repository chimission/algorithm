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
	header.Print()
	if header.Len() != 6 {
		t.Errorf("LinkNode Add error, Get %d, Should %d", header.Len(), 6)
	}
}

func TestLinkDelete(t *testing.T) {
	header := NewLinkNode(1)
	header.Add(2)
	header.Add(3)
	header.Add(4)
	header.Add(5)
	header.Add(6)
	header = header.Delete(1)
	header.Print()
	if header.Len() != 5 {
		t.Errorf("LinkNode Delete error, Get %d, Should %d", header.Len(), 5)
	}
	header = header.Delete(0)
	header.Print()
	if header.Len() != 4 {
		t.Errorf("LinkNode Delete error, Get %d, Should %d", header.Len(), 4)
	}
	header = header.Delete(10) //PASS NO ERROR
	header.Print()
	if header.Len() != 4 {
		t.Errorf("LinkNode Delete error, Get %d, Should %d", header.Len(), 4)
	}
}

func TestLinkInsert(t *testing.T) {
	header := NewLinkNode(1)
	header.Insert(2, 1)
	header.Insert(3, 2)
	header.Insert(4, 3)
	header.Insert(5, 4)
	header.Insert(6, 5)
	header.Print()
	if header.Len() != 6 {
		t.Errorf("LinkNode Insert error, Get %d, Should %d", header.Len(), 6)
	}
	header = header.Insert(0, 0)
	header.Print()
	if header.Len() != 7 {
		t.Errorf("LinkNode Insert error, Get %d, Should %d", header.Len(), 7)
	}
}
