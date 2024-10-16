package main

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3194 lang=golang
 * @lcpr version=20001
 *
 * [3194] 最小元素和最大元素的最小平均值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumAverage(nums []int) float64 {
	res := math.MaxFloat64
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		res = min(res, float64(nums[i]+nums[j])/2)
		i++
		j--
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [7,8,3,4,15,13,4,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,9,8,3,10,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,7,8,9]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		want float64
	}{
		{nums: []int{7, 8, 3, 4, 15, 13, 4, 1}, want: 5.5},
		{nums: []int{1, 9, 8, 3, 10, 5}, want: 5.5},
		{nums: []int{1, 2, 3, 7, 8, 9}, want: 5.0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, minimumAverage(tt.nums))
	}
}
