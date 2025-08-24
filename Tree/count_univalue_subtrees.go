// https://leetcode.com/problems/count-univalue-subtrees/editorial/

package tree

func dfsCount(root *TreeNode, count *int) bool {
	if root == nil {
		return true
	}

	isLeftUnival := dfsCount(root.Left, count)
	isRightUnival := dfsCount(root.Right, count)

	if isLeftUnival && isRightUnival &&
		(root.Left == nil || root.Left.Val == root.Val) &&
		(root.Right == nil || root.Right.Val == root.Val) {
		*count++
		return true
	}

	return false
}

func countUnivalSubtrees(root *TreeNode) int {
	count := 0
	dfsCount(root, &count)
	return count
}
