package tree

import (
	"slices"
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

var traversedTree = &Node{
	Val: 1,
	Left: &Node{
		Val:   2,
		Left:  &Node{Val: 4},
		Right: &Node{Val: 5},
	},
	Right: &Node{
		Val:   3,
		Left:  &Node{Val: 6},
		Right: &Node{Val: 7},
	},
}

func TestInOrderTraversal(t *testing.T) {
	result := traversal(traversedTree, InOrder)
	if !slices.Equal(result, []int{4, 2, 5, 1, 6, 3, 7}) {
		t.Errorf("incorrect traversal")
	}
}

func TestPreOrderTraversal(t *testing.T) {
	result := traversal(traversedTree, PreOrder)
	if !slices.Equal(result, []int{1, 2, 4, 5, 3, 6, 7}) {
		t.Errorf("incorrect traversal")
	}
}

func TestPostOrderTraversal(t *testing.T) {
	result := traversal(traversedTree, PostOrder)
	if !slices.Equal(result, []int{4, 5, 2, 6, 7, 3, 1}) {
		t.Errorf("incorrect traversal")
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	result := traversal(traversedTree, LevelOrder)
	if !slices.Equal(result, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("incorrect traversal")
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

func TestInvertTree(t *testing.T) {
	original := BuildTree([]int{1, 2, 3})
	leftBefore := original.Left.Val
	inverted := invertTree(original)
	rightAfter := inverted.Right.Val

	if leftBefore != rightAfter {
		t.Errorf("Incorrectly inverted")
	}
}

func TestIsSymmetric(t *testing.T) {
	var tree *Node
	tree = BuildTree([]int{1})
	if !isSymmetric(tree) {
		t.Errorf("should be symmetric")
	}

	tree = BuildTree([]int{1, 2, 2, 3, 4, 4, 3})

	if !isSymmetricRec(tree) {
		t.Errorf("should be symmetric")
	}

	tree = BuildTree([]int{1, 2, 2, -1, 3, -1, 3})
	if isSymmetric(tree) {
		t.Errorf("shouldn't be symmetric")
	}
}

func TestAverageOfLevels(t *testing.T) {
	var tree *Node

	tree = BuildTree([]int{1, 2, 2})
	if !slices.Equal(averageOfLevels(tree), []float64{1, 2}) {
		t.Errorf("should be symmetric")
	}

	tree = BuildTree([]int{1, 2, 2, 3, 4, 4, 3})
	result := averageOfLevels(tree)
	if !slices.Equal(result, []float64{1, 2, 3.5}) {
		t.Errorf("incorrect averages")
	}

	tree = BuildTree([]int{3, 9, 20, -1, -1, 15, 7})
	if !slices.Equal(averageOfLevels(tree), []float64{3, 14.5, 11}) {
		t.Errorf("incorrect averages")
	}
}
