package t5

import "testing"

/**
dp[i][j] = dp[i+1][j-1]
j-1 <= i+1
j - i <= 2
*/
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	// 进行初始化
	dp, begin, maxLen := make([][]bool, n), 0, 0
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for j := 1; j < n; j++ {
		for i := 0; i <= j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j - i <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && (j-i+1) > maxLen {
				maxLen = j-i+1
				begin=i
			}
		}
	}
	return s[begin:begin+maxLen]
}

func TestDo5(t *testing.T)  {
	s := longestPalindrome("babad")
	print(s)
}
