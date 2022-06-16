package structure

import (
	"testing"
)

func TestDoubleLinkInit(t *testing.T) {
	doubleLink := InitDoubleLink(1)
	if (doubleLink.head != doubleLink.tail) && (doubleLink.head != nil) {
		t.Error("new double link error")
	}
}
