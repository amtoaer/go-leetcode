package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	var (
		res   []string
		start int
	)
	for i := 0; i < len(nums); i++ {
		if i+1 == len(nums) || nums[i+1] != nums[i]+1 {
			if start == i {
				res = append(res, strconv.Itoa(nums[i]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
			}
			start = i + 1
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, []string{"0->2", "4->5", "7"}, summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	assert.Equal(t, []string{"0", "2->4", "6", "8->9"}, summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
