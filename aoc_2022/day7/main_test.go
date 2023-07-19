package main

import "testing"

func TestTreeCreation(t *testing.T) {
	root := readInput("./test_input.txt")
	if root == nil {
		t.Errorf("expected file tree")
	}

	if root.Children["d"] == nil {
		t.Errorf("expected directory under root > d")
	}

	if root.Children["a"].Children["f"] == nil {
		t.Errorf("expected file under root > a > f")
	}
}

func TestDirectorySize(t *testing.T) {
	root := readInput("./test_input.txt")
	if root == nil {
		t.Errorf("expected file tree")
	}

	calculateSize(root)
	size := findSize(root)
	if size != 95437 {
		t.Errorf("expected size to be 95437, but got %d", size)
	}
}
