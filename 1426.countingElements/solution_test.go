/*
 * @lc app=leetcode.cn id=1426 lang=golang
 * @lcpr version=30202
 *
 * [1426] 数元素
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func countElements(arr []int) int {
	m := make(map[int]int)
	for _, item := range arr {
		m[item-1]++
	}
	var res int
	for _, item := range arr {
		if _, ok := m[item]; ok {
			res++
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 1, 3, 3, 5, 5, 7, 7}, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, countElements(tt.arr))
	}
}

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,3,3,5,5,7,7]\n
// @lcpr case=end

*/
