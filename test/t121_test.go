package test

import (
	"math"
	"testing"
)

func maxProfit3(prices []int) int {
	n := len(prices)
	min, max := prices[0], prices[n-1]
	for min < max {

	}
	return max - min
}

func maxProfit2(prices []int) int {
	n := len(prices)
	min, max := prices[0], prices[n-1]
	for p, q := 0, n-1; p <= q; {
		if min-prices[p] > prices[q]-max {
			if prices[p] < min {
				min = prices[p]
			}
			p++
		} else {
			if prices[q] > max {
				max = prices[q]
			}
			q--
		}
	}
	return max - min
}

func maxProfit(prices []int) int {
	max, min := 0, math.MaxInt
	for i, n := 0, len(prices); i < n; i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}

func TestDo121(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(prices))
}
