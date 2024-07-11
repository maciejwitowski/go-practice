package tree

// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

/*
BuildTree Build the tree based on the array representation ("heap numbering")

For any node at index i:
  - Its left child is at index 2i + 1
  - Its right child is at index 2i + 2

Conversely, for any child at index j, its parent is at index (j-1)/2 (integer division)
*/
func BuildTree(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	nodes := make([]*Node, len(nums))
	for i, num := range nums {
		if num != -1 {
			nodes[i] = &Node{Val: num}
		}
	}

	for i := range nodes {
		if nodes[i] == nil {
			continue
		}

		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		if leftIndex < len(nodes) {
			nodes[i].Left = nodes[leftIndex]
		}
		if rightIndex < len(nodes) {
			nodes[i].Right = nodes[rightIndex]
		}
	}

	return nodes[0]
}

// MaxDepth Calculate max depth of a tree
func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

func IsBalanced(root *Node) bool {
	_, balanced := checkBalance(root)
	return balanced
}

func checkBalance(root *Node) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftBalanced := checkBalance(root.Left)
	rightDepth, rightBalanced := checkBalance(root.Right)

	depth := max(leftDepth, rightDepth) + 1
	isBalanced := leftBalanced && rightBalanced && abs(leftDepth-rightDepth) <= 1

	return depth, isBalanced
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
