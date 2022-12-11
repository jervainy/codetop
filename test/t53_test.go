package test

import "math"

func maxSubArray0(nums []int) int {
	ans, previous := nums[0], nums[0]
	for i, n := 1, len(nums); i < n; i++ {
		previous = int(math.Max(float64(nums[i]+previous), float64(nums[i])))
		ans = int(math.Max(float64(previous), float64(ans)))
	}
	return ans
}

func maxSubArray(nums []int) int {
	dp, max := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i, n := 1, len(nums); i < n; i++ {
		if dp[i-1] >= 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
