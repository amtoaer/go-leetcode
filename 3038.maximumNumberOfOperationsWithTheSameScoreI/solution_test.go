package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3038 lang=golang
 *
 * [3038] Maximum Number of Operations With the Same Score I
 */

// @lc code=start
func maxOperations(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sum, count := nums[0]+nums[1], 1
	for i := 2; i+1 < len(nums); i += 2 {
		if nums[i]+nums[i+1] != sum {
			break
		}
		count += 1
	}
	return count
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		nums   []int
		output int
	}{
		{[]int{3, 2, 1, 4, 5}, 2},
		{[]int{3, 2, 6, 1, 4}, 1},
	}
	for _, testcase := range testcases {
		assert.Equal(t, testcase.output, maxOperations(testcase.nums))
	}
}
