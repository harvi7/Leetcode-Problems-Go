package backtracking

func permute(nums []int) [][]int {
	var res [][]int
	permuteHelper(nums, []int{}, &res)
	return res
}

func permuteHelper(source []int, ans []int, ansList *[][]int) {
	if len(source) == 0 {
		*ansList = append(*ansList, ans)
		return
	}
	for i, e := range source {
		newSource := append(append([]int{}, source[:i]...), source[i+1:]...)
		permuteHelper(newSource, append(ans, e), ansList)
	}
}
