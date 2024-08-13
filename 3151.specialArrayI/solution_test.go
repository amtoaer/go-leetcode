package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3151 lang=golang
 *
 * [3151] Special Array I
 */

// @lc code=start
func isArraySpecial(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1]&1 == nums[i]&1 {
			return false
		}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output bool
	}{
		{[]int{1}, true},
		{[]int{2, 1, 4}, true},
		{[]int{4, 3, 1, 6}, false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, isArraySpecial(tt.input))
	}
}
