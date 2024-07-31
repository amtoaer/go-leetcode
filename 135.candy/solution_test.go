package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] Candy
 */

// @lc code=start
func candy(ratings []int) int {
	left := make([]int, len(ratings))
	left[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 1
	sum := max(right, left[len(ratings)-1])
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		sum += max(right, left[i])
	}
	return sum
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{input: []int{1, 0, 2}, output: 5},
		{input: []int{1, 2, 2}, output: 4},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, candy(tt.input))
	}
}
