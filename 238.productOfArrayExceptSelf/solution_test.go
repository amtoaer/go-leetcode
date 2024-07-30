package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	tmp := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = answer[i] * tmp
		tmp *= nums[i]
	}
	return answer
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, productExceptSelf(tt.input))
	}
}
