package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	start := 0
	for idx := 0; idx < len(nums); idx++ {
		if nums[idx] == val {
			continue
		}
		nums[start] = nums[idx]
		start++
	}
	return start
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 2, removeElement([]int{3, 2, 2, 3}, 3))
}
