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

func isSameTree(p *Node, q *Node) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func invertTree(root *Node) *Node {
	if root == nil {
		return root
	}

	temp := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(temp)
	return root
}

const (
	InOrder = iota
	PreOrder
	PostOrder
	LevelOrder
)

func traversal(root *Node, technique int) []int {
	var result []int
	switch technique {
	case InOrder:
		inOrderTraversal(root, &result)
	case PreOrder:
		preOrderTraversal(root, &result)
	case PostOrder:
		postOrderTraversal(root, &result)
	case LevelOrder:
		levelOrderTraversal(root, &result)
	}
	return result
}

// left -> root -> right
func inOrderTraversal(root *Node, result *[]int) {
	if root == nil {
		return
	}

	inOrderTraversal(root.Left, result)
	*result = append(*result, root.Val)
	inOrderTraversal(root.Right, result)
}

// left -> right -> root
func postOrderTraversal(root *Node, result *[]int) {
	if root == nil {
		return
	}

	postOrderTraversal(root.Left, result)
	postOrderTraversal(root.Right, result)
	*result = append(*result, root.Val)
}

// root -> left -> right
func preOrderTraversal(root *Node, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	preOrderTraversal(root.Left, result)
	preOrderTraversal(root.Right, result)
}

func levelOrderTraversal(root *Node, result *[]int) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		*result = append(*result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}

	queue := []*Node{root.Left, root.Right}

	for len(queue) > 0 {
		left, right := queue[0], queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}

	return true
}

func isSymmetricRec(root *Node) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *Node) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val &&
		isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}

func averageOfLevels(root *Node) []float64 {
	if root == nil {
		return []float64{}
	}

	var result []float64
	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, float64(levelSum)/float64(levelSize))
	}
	return result
}
