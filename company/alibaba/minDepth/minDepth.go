package minDepth

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	} else {
		return int(math.Min(float64(left), float64(right))) + 1
	}
}

// bfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	q := make([]*TreeNode, 1)
	q[0] = root
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		depth++
	}
	return depth
}
