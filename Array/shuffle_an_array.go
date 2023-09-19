type Solution struct {
    original []int
}


func Constructor(nums []int) Solution {
    return Solution{
        original: nums,
    }
}


func (this *Solution) Reset() []int {
    return this.original
}


func (this *Solution) Shuffle() []int {
    shuffled := make([]int, len(this.original))
    copy(shuffled, this.original)
    for i := 0; i< len(shuffled); i++ {
        r := rand.Intn(len(shuffled))
        shuffled[i], shuffled[r] = shuffled[r], shuffled[i]
    }
    return shuffled
}