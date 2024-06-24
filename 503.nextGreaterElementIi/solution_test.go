package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] Next Greater Element II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	for i := range res {
		res[i] = -1
	}
	var stack []int
	// 对于循环数组，可以重复两次并对 index 取模达到循环效果
	for i := 0; i < l*2-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%l] {
			res[stack[len(stack)-1]] = nums[i%l]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%l)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		in  []int
		out []int
	}{
		{in: []int{1, 2, 1}, out: []int{2, -1, 2}},
		{in: []int{1, 2, 3, 4, 3}, out: []int{2, 3, 4, -1, 4}},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.out, nextGreaterElements(tc.in))
	}
}
