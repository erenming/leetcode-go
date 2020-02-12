package isBalanced

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution 1
// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
// }
//
// func height(node *TreeNode) int {
// 	if node == nil {
// 		return 0
// 	}
// 	return int(math.Max(float64(height(node.Left)), float64(height(node.Right)))) + 1
// }

// Solution 2
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	balanced := true
	height(root, &balanced)
	return balanced
}

func height(root *TreeNode, b *bool) int {
	if root == nil {
		return 0
	}
	left := height(root.Left, b)
	right := height(root.Right, b)
	if math.Abs(float64(left - right)) > 1 {
		*b = false
		return -1
	}
	return int(math.Max(float64(left), float64(right))) + 1
}