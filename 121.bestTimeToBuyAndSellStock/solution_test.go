package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	var cur, res int
	for i := 1; i < len(prices); i++ {
		cur = max(prices[i]-prices[i-1]+cur, 0)
		res = max(res, cur)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, maxProfit(tt.input))
	}
}
