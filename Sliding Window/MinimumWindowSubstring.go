# https://www.youtube.com/watch?v=jSto0O4AJbM

func minWindow(s string, t string) string {
    countT := map[uint8]int{}
    for i, _ := range t {
        countT[t[i]]++
    }
    l, r := 0, 0
    window := map[uint8]int{}
    res := [2]int{0, math.MaxInt32}
    have := 0
    for r < len(s) {
        if countT[s[r]] > window[s[r]] {
            have++
        }
        window[s[r]]++
        for l <= r && window[s[l]] > countT[s[l]] {
            window[s[l]]--
            l++
        }
        if have == len(t) && r - l + 1 < res[1] - res[0] + 1{
            res[0], res[1] = l, r
        }
        r++
    }
    if res[1] == math.MaxInt32 {
        return ""
    }
    return s[res[0] : res[1] + 1]
}