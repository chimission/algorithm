package sort

import "testing"

var unsorted = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var sorted = []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}

func TestInsertSort(t *testing.T) {
	unsorted_copy := MakeCopy(unsorted)
	sored_list := InsertSort(unsorted_copy)
	if !LoopCompare(sored_list, sorted) {
		t.Errorf("sorted %v", unsorted)
		t.Errorf("   got %v", sored_list)
	}
}

func TestMerge(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 10}
	b := []int{2, 4, 6, 8}
	sored_list := Merge(a, b)
	corret_sored_list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !LoopCompare(sored_list, corret_sored_list) {
		t.Errorf("merge %v", corret_sored_list)
		t.Errorf("   got %v", sored_list)
	}
}

func TestMergeSort(t *testing.T) {
	unsorted_copy := MakeCopy(unsorted)
	sored_list := InsertSort(unsorted_copy)
	if !LoopCompare(sored_list, sorted) {
		t.Errorf("sorted %v", unsorted)
		t.Errorf("   got %v", sored_list)
	}
}

func TestBubbleSort(t *testing.T) {
	unsorted_copy := MakeCopy(unsorted)
	sored_list := BubbleSort(unsorted_copy)
	if !LoopCompare(sored_list, sorted) {
		t.Errorf("sorted %v", unsorted)
		t.Errorf("   got %v", sored_list)
	}
}
