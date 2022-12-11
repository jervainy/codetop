package test

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, res := []*TreeNode{root}, make([][]int, 0)
	for level := 0; len(queue) > 0; level++ {
		q, vals := queue, make([]int, 0)
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level%2 == 1 {
			for left, right := 0, len(vals)-1; left < right; {
				vals[left], vals[right] = vals[right], vals[left]
				left++
				right--
			}
		}
		res = append(res, vals)
	}
	return res
}

func TestDo103(t *testing.T) {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	res := zigzagLevelOrder(root)
	println(res)
}
