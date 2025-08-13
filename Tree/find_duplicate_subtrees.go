package tree

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	tripletToID := make(map[string]int)
	cnt := make(map[int]int)
	traverse(root, tripletToID, cnt, &res)
	fmt.Print(tripletToID)
	return res
}

func traverse(node *TreeNode, tripletToID map[string]int, cnt map[int]int, res *[]*TreeNode) int {
	if node == nil {
		return 0
	}

	triplet := fmt.Sprintf("%d,%d,%d",
		traverse(node.Left, tripletToID, cnt, res),
		node.Val,
		traverse(node.Right, tripletToID, cnt, res),
	)

	if _, exists := tripletToID[triplet]; !exists {
		tripletToID[triplet] = len(tripletToID) + 1
	}

	id := tripletToID[triplet]
	cnt[id]++
	if cnt[id] == 2 {
		*res = append(*res, node)
	}

	return id
}
