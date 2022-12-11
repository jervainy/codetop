package test

import "testing"

func twoSum(nums []int, target int) []int {
	m, n := make(map[int]int), len(nums)
	m[nums[0]] = 0
	for i := 1; i < n; i++ {
		if index, ok := m[target-nums[i]]; ok {
			return []int{index, i}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

func TestDo1(t *testing.T) {

}
