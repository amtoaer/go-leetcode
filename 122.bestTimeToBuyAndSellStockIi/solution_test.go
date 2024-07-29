package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	var cur int
	for i := 1; i < len(prices); i++ {
		cur = max(prices[i]-prices[i-1]+cur, cur)
	}
	return cur
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{input: []int{7, 1, 5, 3, 6, 4}, output: 7},
		{input: []int{1, 2, 3, 4, 5}, output: 4},
		{input: []int{7, 6, 4, 3, 1}, output: 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, maxProfit(tt.input))
	}
}
