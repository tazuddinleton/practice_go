package main

import "testing"

func TestSum(t *testing.T) {
	sum := getSum("./test_input.txt")

	if sum != 157 {
		t.Errorf("expected 157, got %d", sum)
	}
}
