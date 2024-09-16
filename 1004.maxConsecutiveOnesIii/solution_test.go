package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	var res, left, cntK int
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			cntK++
		}
		for cntK > k {
			if nums[left] == 0 {
				cntK--
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, longestOnes(tt.nums, tt.k))
	}
}
