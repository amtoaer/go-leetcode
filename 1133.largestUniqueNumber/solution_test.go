/*
 * @lc app=leetcode.cn id=1133 lang=golang
 * @lcpr version=30202
 *
 * [1133] 最大唯一数
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func largestUniqueNumber(nums []int) int {
	m := make(map[int]int)
	res := -1
	for _, num := range nums {
		m[num]++
	}
	for val, cnt := range m {
		if cnt == 1 {
			res = max(res, val)
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		want int
	}{
		{[]int{5, 7, 3, 9, 4, 9, 8, 3, 1}, 8},
		{[]int{9, 9, 8, 8}, -1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, largestUniqueNumber(tt.nums))
	}
}

/*
// @lcpr case=start
// [5,7,3,9,4,9,8,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,8,8]\n
// @lcpr case=end

*/
