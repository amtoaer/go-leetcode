/*
 * @lc app=leetcode.cn id=163 lang=golang
 * @lcpr version=30202
 *
 * [163] 缺失的区间
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func findMissingRanges(nums []int, lower int, upper int) [][]int {
	res := [][]int{}
	left := lower
	for _, right := range nums {
		if right > left {
			res = append(res, []int{left, right - 1})
		}
		left = right + 1
	}
	if left <= upper {
		res = append(res, []int{left, upper})
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums  []int
		lower int
		upper int
		want  [][]int
	}{
		{[]int{0, 1, 3, 50, 75}, 0, 99, [][]int{{2, 2}, {4, 49}, {51, 74}, {76, 99}}},
		{[]int{-1}, -1, -1, [][]int{}},
		{[]int{}, -3, -1, [][]int{{-3, -1}}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findMissingRanges(tt.nums, tt.lower, tt.upper))
	}
}

/*
// @lcpr case=start
// [0, 1, 3, 50, 75]\n0\n99\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n-1\n-1\n
// @lcpr case=end

*/
