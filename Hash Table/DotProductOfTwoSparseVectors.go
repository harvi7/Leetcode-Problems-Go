package hashtable

// https://www.youtube.com/watch?v=7OWu2_tTBN4

type SparseVector struct {
	nonZeroes map[int]int
}

func Constructor(nums []int) SparseVector {
	nonZeroes := map[int]int{}

	for index, num := range nums {
		if num != 0 {
			nonZeroes[index] = num
		}
	}

	return SparseVector{nonZeroes}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	result := 0

	// go through all of the non zero indices in this
	for index, num := range this.nonZeroes {
		// check if also in other's cache
		if otherNum, ok := vec.nonZeroes[index]; ok {
			result += num * otherNum
		}
	}

	return result
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
