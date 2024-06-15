package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2779 lang=golang
 *
 * [2779] Maximum Beauty of an Array After Applying Operation
 */

// @lc code=start
func max(nums ...int) int {
	res := math.MinInt
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}

func min(nums ...int) int {
	res := math.MaxInt
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}

func maximumBeauty(nums []int, k int) int {
	counts := max(nums...)
	diff := make([]int, counts+2)
	for _, num := range nums {
		diff[max(0, num-k-1)]++
		diff[min(counts+1, num+k)]--
	}
	cur, res := 0, 0
	for _, x := range diff {
		cur += x
		res = max(res, cur)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{4, 6, 1, 2}, 2, 3},
		{[]int{1, 1, 1, 1}, 10, 4},
		{[]int{52, 34}, 21, 2},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, maximumBeauty(tc.nums, tc.k))
	}
}
