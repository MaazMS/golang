package main

import "testing"

func TestYourTest(t *testing.T) {
	type test struct {
		data []int
		ans  int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
	}

	for _, v := range tests {
		x := yoursum(v.data...)
		if x != v.ans {
			t.Error("Expected", v.ans, "Got", x)
		}
	}
}
