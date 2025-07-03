package tree

func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftTail := dfs(root.Left)
	rightTail := dfs(root.Right)

	if root.Left != nil {
		leftTail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	if leftTail == nil && rightTail == nil {
		return root
	} else if rightTail == nil {
		return leftTail
	} else {
		return rightTail
	}
}
