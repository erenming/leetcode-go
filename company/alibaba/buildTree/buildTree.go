package buildTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || postorder == nil || len(postorder) != len(inorder) {
		return nil
	}
	mm := make(map[int]int)
	for i, v := range inorder {
		mm[v] = i
	}
	return dfs(0, len(inorder)-1, 0, len(postorder)-1, postorder, mm)
}

func dfs(is, ie, ps, pe int, postorder []int, mm map[int]int) *TreeNode {
	if pe < ps || ie < is {
		return nil
	}
	root := &TreeNode{postorder[pe], nil, nil}
	idx, _ := mm[postorder[pe]]
	leftlen, rightlen := (idx-1)-is+1, ie-(idx+1)+1
	left := dfs(is, idx-1, ps, pe-1-rightlen, postorder, mm)
	right := dfs(idx+1, ie, ps+leftlen, pe-1, postorder, mm)
	root.Left = left
	root.Right = right
	return root
}
