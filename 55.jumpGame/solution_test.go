package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	var farthest int
	for idx, num := range nums {
		if idx > farthest {
			break
		}
		farthest = max(farthest, idx+num)
	}
	return farthest >= len(nums)-1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		expect bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.expect, canJump(tt.input))
	}
}
