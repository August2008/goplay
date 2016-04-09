package arrays

import (
	"testing"
)

func TestCanAddNumbers(t *testing.T) {
	var arr = []int{1, 2, 3}
	var result = Add(arr...)
	if result != 6 {
		t.Log("Ohh, crap, I failed!")
		t.Fail()
	}
}
