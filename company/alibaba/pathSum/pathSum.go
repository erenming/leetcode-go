package pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	dfs(root, sum, make([]int, 0), &res)
	return res
}

func dfs(node *TreeNode, sum int, cur []int, res *[][]int) {
	if node.Left == nil && node.Right == nil {
		if sum == node.Val {
			*res = append(*res, append(cur, node.Val))
		}
		return
	}
	c := append(append([]int{}, cur...), node.Val)
	if node.Left != nil {
		dfs(node.Left, sum-node.Val, c, res)
	}
	if node.Right != nil {
		dfs(node.Right, sum-node.Val, c, res)
	}
}
