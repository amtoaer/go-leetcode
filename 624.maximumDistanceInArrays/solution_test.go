/*
 * @lc app=leetcode.cn id=624 lang=golang
 * @lcpr version=30202
 *
 * [624] 数组列表中的最大距离
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func maxDistance(arrays [][]int) int {
	var res int
	if len(arrays) == 0 {
		return res
	}
	arr := arrays[0]
	minVal, maxVal := arr[0], arr[len(arr)-1]
	for _, arr := range arrays[1:] {
		tmpMin, tmpMax := arr[0], arr[len(arr)-1]
		res = max(res, maxVal-tmpMin, tmpMax-minVal)
		minVal = min(minVal, tmpMin)
		maxVal = max(maxVal, tmpMax)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		arrays [][]int
		want   int
	}{
		{[][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}, 4},
		{[][]int{{1}, {1}}, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, maxDistance(tt.arrays))
	}
}

/*
// @lcpr case=start
// [[1,2,3],[4,5],[1,2,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[1],[1]]\n
// @lcpr case=end

*/
