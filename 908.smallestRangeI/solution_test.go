package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=908 lang=golang
 * @lcpr version=20002
 *
 * [908] 最小差值 I
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func smallestRangeI(nums []int, k int) int {
	maxVal, minVal := math.MinInt, math.MaxInt
	for _, num := range nums {
		maxVal = max(maxVal, num)
		minVal = min(minVal, num)
	}
	return max(maxVal-minVal-2*k, 0)
}

// @lc code=end

/*
// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0,10]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,3,6]\n3\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1}, k: 0, want: 0},
		{nums: []int{0, 10}, k: 2, want: 6},
		{nums: []int{1, 3, 6}, k: 3, want: 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, smallestRangeI(tt.nums, tt.k))
	}
}
