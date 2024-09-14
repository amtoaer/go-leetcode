package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2848 lang=golang
 *
 * [2848] Points That Intersect With Cars
 */

// @lc code=start
func numberOfPoints(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})
	var (
		rg  [2]int = [2]int{nums[0][0], nums[0][1]}
		res int
	)
	for i := 0; i < len(nums); i++ {
		if nums[i][0] > rg[1] {
			res += rg[1] - rg[0] + 1
			rg = [2]int{nums[i][0], nums[i][1]}
		} else {
			rg[1] = max(rg[1], nums[i][1])
		}
	}
	res += rg[1] - rg[0] + 1
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums [][]int
		want int
	}{
		{nums: [][]int{{3, 6}, {1, 5}, {4, 7}}, want: 7},
		{nums: [][]int{{1, 3}, {5, 8}}, want: 7},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, numberOfPoints(tt.nums))
	}
}
