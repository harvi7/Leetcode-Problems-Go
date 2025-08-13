package tree

func checkEqualTree(root *TreeNode) bool {
	var seen []int

	// helper function to calculate subtree sums
	var sum func(node *TreeNode) int
	sum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		val := sum(node.Left) + sum(node.Right) + node.Val
		seen = append(seen, val)
		return val
	}

	total := sum(root)
	// pop the last element (the sum of the whole tree)
	seen = seen[:len(seen)-1]

	if total%2 == 0 {
		half := total / 2
		for _, s := range seen {
			if s == half {
				return true
			}
		}
	}
	return false
}
