package main

import "testing"

func TestTreeCreation(t *testing.T) {
	root := readInput("./test_input.txt")
	if root == nil {
		t.Errorf("expected file tree")
	}

	if root.Children["cvt"] == nil {
		t.Errorf("expected directory under root > cvt")
	}

	if root.Children["cvt"].Children["file.txt"] == nil {
		t.Errorf("expected file under root > cvt > file.txt")
	}
}
