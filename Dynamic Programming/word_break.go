func wordBreak(s string, wordDict []string) bool {
    if wordDict == nil || len(wordDict) == 0 { return false }
    dp := make([]bool, len(s) + 1)
    dp[0] = true
    sort.Strings(wordDict)

    for end := 1; end <= len(s); end++ {
        for i := 0; i < end; i++ {
            if dp[i] && stringInSlice(s[i : end], wordDict) {
                dp[end] = true
                break
            }
        }
    }
    return dp[len(s)]
}

 func stringInSlice(str string, list []string) bool {
 	for _, v := range list {
 		if v == str {
 			return true
 		}
 	}
 	return false
 }