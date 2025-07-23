/*
 * @lc app=leetcode.cn id=739 lang=golang
 * @lcpr version=30202
 *
 * [739] 每日温度
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	var stack []int
	res := make([]int, len(temperatures))
	for idx, templtemperature := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < templtemperature {
			prevIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prevIdx] = idx - prevIdx
		}
		stack = append(stack, idx)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		temperatures []int
		want         []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, dailyTemperatures(tt.temperatures))
	}
}

/*
// @lcpr case=start
// [73,74,75,71,69,72,76,73]\n
// @lcpr case=end

// @lcpr case=start
// [30,40,50,60]\n
// @lcpr case=end

// @lcpr case=start
// [30,60,90]\n
// @lcpr case=end

*/
