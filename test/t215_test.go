package test

import (
	"testing"
)

// ---------------------
func quickSort(nums []int, left, right int) {
	if left < right {
		partIndex := partition(nums, left, right)
		quickSort(nums, left, partIndex-1)
		quickSort(nums, partIndex+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if nums[i] > nums[pivot] {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	nums[index-1], nums[pivot] = nums[pivot], nums[index-1]
	return index - 1
}

func findKthLargestQuickSort(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[k-1]
}

func TestDo215(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	result := findKthLargestQuickSort(nums, 4)
	println(result)
}

// ----------------------------------------------------
// 使用最小堆
func findKthLargestHeap(nums []int, k int) int {
	buildMinHeap(nums, k)
	for i, n := k, len(nums); i < n; i++ {
		if nums[i] > nums[0] {
			nums[i], nums[0] = nums[0], nums[i]
			shiftDown(nums, 0, k)
		}
	}
	return nums[0]
}

func buildMinHeap(nums []int, heapSize int) {
	for i := 1; i < heapSize; i++ {
		shiftUp(nums, i)
	}
}

// 进行堆的上升
func shiftUp(nums []int, index int) {
	parent := (index - 1) / 2
	if index > 0 && index > parent && nums[index] < nums[parent] {
		nums[index], nums[parent] = nums[parent], nums[index]
		shiftUp(nums, parent)
	}
}

// 进行堆的下沉
func shiftDown(nums []int, index, heapSize int) {
	left, right, leaf := 2*index+1, 2*index+2, index
	if left < heapSize && nums[index] > nums[left] {
		leaf = left
	}
	if right < heapSize && nums[index] > nums[right] {
		if leaf == index {
			leaf = right
		} else if nums[right] < nums[left] {
			leaf = right
		}
	}
	if leaf != index {
		nums[index], nums[leaf] = nums[leaf], nums[index]
		shiftDown(nums, leaf, heapSize)
	}
}
