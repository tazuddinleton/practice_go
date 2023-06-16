package main

import "testing"

func TestOverlap(t *testing.T) {
	overlap, _ := findOverlaps("./test_input.txt")

	if overlap != 4 {
		t.Errorf("wanted 4, got %d", overlap)
	}
}
