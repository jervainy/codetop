package test

import "testing"

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, left, right int) int {
	if right-left < 0 {
		return -1
	}
	mid := (right-left)/2 + left
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target && target >= nums[0] {
		return binarySearch(nums, target, left, mid-1)
	} else if nums[mid] < nums[0] && target >= nums[0] {
		return binarySearch(nums, target, left, mid-1)
	} else if nums[mid] < nums[0] && target < nums[mid] {
		return binarySearch(nums, target, left, mid-1)
	} else {
		return binarySearch(nums, target, mid+1, right)
	}
}

func TestDo33(t *testing.T) {
	nums, target := []int{4, 5, 6, 7, 0, 1, 2}, 0
	res := search(nums, target)
	println(res)
}
