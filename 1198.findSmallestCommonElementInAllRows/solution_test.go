/*
 * @lc app=leetcode.cn id=1198 lang=golang
 * @lcpr version=30202
 *
 * [1198] 找出所有行中最小公共元素
 */

package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func smallestCommonElement(mat [][]int) int {
	m := make(map[int]int)
	for _, row := range mat {
		for _, val := range row {
			m[val]++
		}
	}
	res := math.MaxInt
	for num, cnt := range m {
		if cnt == len(mat) {
			res = min(res, num)
		}
	}
	if res == math.MaxInt {
		res = -1
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		mat  [][]int
		want int
	}{
		{[][]int{{1, 2, 3, 4, 5}, {2, 4, 5, 8, 10}, {3, 5, 7, 9, 11}, {1, 3, 5, 7, 9}}, 5},
		{[][]int{{1, 2, 3}, {2, 3, 4}, {2, 3, 5}}, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, smallestCommonElement(tt.mat))
	}
}

/*
// @lcpr case=start
// [[1,2,3,4,5],[2,4,5,8,10],[3,5,7,9,11],[1,3,5,7,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[2,3,4],[2,3,5]]\n
// @lcpr case=end

*/
