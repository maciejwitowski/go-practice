package tree

import (
	"testing"
)

/*
		3
		/  \
	  9 	20
		    /    \
	       15   7
*/
func setupTest() *Node {
	input := []int{3, 9, 20, -1, -1, 15, 7}
	return BuildTree(input)
}

func TestBuildsTree(t *testing.T) {
	tree := setupTest()

	if tree.Right.Left.Val != 15 {
		t.Errorf("Expected 15, got %d", tree.Right.Left.Val)
	}
}

func TestMaximumDepth(t *testing.T) {
	tree := setupTest()
	depth := MaxDepth(tree)
	if depth != 3 {
		t.Errorf("Expected 3, got %d", depth)
	}

	// Empty tree
	tree = BuildTree([]int{})
	depth = MaxDepth(tree)
	if depth != 0 {
		t.Errorf("Expected 0, got %d", depth)
	}

	// Empty tree
	tree = setupTest()
	depth = MaxDepth(tree)
	if depth != 3 {
		t.Errorf("Expected 3, got %d", depth)
	}
}

func TestIsBalanced(t *testing.T) {
	tree := setupTest()
	if !IsBalanced(tree) {
		t.Errorf("Expected a balanced tree")
	}
}

func TestIsSameTree(t *testing.T) {
	first := setupTest()
	second := setupTest()

	if !isSameTree(first, second) {
		t.Errorf("expected trees to be identical")
	}

	first = BuildTree([]int{1, 2, 1})
	second = BuildTree([]int{1, 1, 2})
	if isSameTree(first, second) {
		t.Errorf("expected trees to be different")
	}
}
