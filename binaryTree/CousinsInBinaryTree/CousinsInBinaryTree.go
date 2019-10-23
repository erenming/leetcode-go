package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var (
		queue []*TreeNode
		xParent *TreeNode
		yParent *TreeNode
	)
	queue = append(queue, root) // push root
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
				if node.Left.Val == x {
					xParent = node
				}
				if node.Left.Val == y {
					yParent = node
				}
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				if node.Right.Val == x {
					xParent = node
				}
				if node.Right.Val == y {
					yParent = node
				}
			}
			length--
			if xParent != nil && yParent != nil {
				break
			}
		}
		if xParent != nil && yParent != nil {
			return xParent != yParent
		}
		if ((xParent != nil && yParent == nil) || (xParent == nil && yParent != nil)) {
			return false
		}
	}
	return false
}
