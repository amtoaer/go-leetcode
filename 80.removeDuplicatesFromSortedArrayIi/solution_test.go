package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	count := 1
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] == nums[start] {
			count++
			if count <= 2 {
				start++
				nums[start] = nums[idx]
			}
		} else {
			start++
			nums[start] = nums[idx]
			count = 1
		}
	}
	return start + 1
}

// @lc code=end

func Test(t *testing.T) {
	res := []int{1, 1, 1, 2, 2, 3}
	assert.Equal(t, 5, removeDuplicates(res))
	assert.Equal(t, []int{1, 1, 2, 2, 3}, res[:5])
}
