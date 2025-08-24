package tree

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	hLeft, hRight := 0, 0
	pLeft, pRight := root, root

	for pLeft != nil {
		hLeft++
		pLeft = pLeft.Left
	}

	for pRight != nil {
		hRight++
		pRight = pRight.Right
	}

	if hLeft == hRight {
		return int(math.Pow(2, float64(hLeft))) - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
