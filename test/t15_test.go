package test

import "testing"

func quickSort15(nums []int, left, right int) {
	if left < right {
		partitionIndex := partition15(nums, left, right)
		quickSort15(nums, left, partitionIndex-1)
		quickSort15(nums, partitionIndex+1, right)
	}
}

func partition15(nums []int, left, right int) int {
	pivot, j := nums[left], left+1
	// [left+1, j) <= pivot
	// [j, i) > pivot
	for i := left + 1; i <= right; i++ {
		if nums[i] <= pivot {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	nums[j-1], nums[left] = nums[left], nums[j-1]
	return j - 1
}

func threeSum(nums []int) [][]int {
	n, res := len(nums), make([][]int, 0)
	quickSort15(nums, 0, n-1)
	if nums[0] > 0 || nums[n-1] < 0 {
		return res
	}
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
				left += 1
				right -= 1
			} else if sum > 0 {
				right -= 1
			} else {
				left += 1
			}
		}
	}
	return res
}

func TestDo15(t *testing.T) {
	nums := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}
	threeSum(nums)
}
