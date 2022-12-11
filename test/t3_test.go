package test

import "testing"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	sl, sm, max, left, right, st := len(s), 0, 0, uint(0), uint(0), make(map[uint8]uint)
	for right >= left && right < uint(sl) {
		c := s[right]
		if rt, ok := st[c]; ok && rt >= left {
			delete(st, c)
			left = rt + 1
			if sm > max {
				max = sm
			}
			sm = int(right - left)
		} else {
			st[c] = right
			sm += 1
			right += 1
		}
	}
	if sm > max {
		max = sm
	}
	return max
}

func TestDo3(t *testing.T) {
	rt := lengthOfLongestSubstring("pwwkew")
	println(rt)
}
