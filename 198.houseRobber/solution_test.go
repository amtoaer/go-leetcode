package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=198 lang=golang
 * @lcpr version=20004
 *
 * [198] 打家劫舍
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(second, first+nums[i])
	}
	return second
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,7,9,3,1]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"example 1", []int{1, 2, 3, 1}, 4},
		{"example 2", []int{2, 7, 9, 3, 1}, 12},
		{"single element", []int{5}, 5},
		{"two elements", []int{1, 2}, 2},
		{"all same", []int{1, 1, 1, 1}, 2},
		{"alternating", []int{2, 1, 2, 1}, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := rob(tc.input)
			if got != tc.expected {
				t.Errorf("rob(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}
