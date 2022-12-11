package test

import "testing"

func isValid(s string) bool {
	var stack []int32
	size, m := 0, map[int32]int32{41: 40, 93: 91, 125: 123}
	for _, c := range s {
		if v, ok := m[c]; ok {
			size--
			if size < 0 { // 出栈
				return false
			}
			if stack[size] != v {
				return false
			}
			stack = stack[:size]
		} else {
			stack = append(stack, c)
			size++
		}
	}
	return size == 0
}

func TestDo20(t *testing.T) {
	println(isValid("()[]{}"))
}
