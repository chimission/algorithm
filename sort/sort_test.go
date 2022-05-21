package sort

import "testing"

func LoopCompare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func MakeCopy(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

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
