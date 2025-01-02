package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=69 lang=golang
 * @lcpr version=20004
 *
 * [69] x 的平方根
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func mySqrt(x int) int {
	var (
		left, right = 0, x
		res         = -1
	)
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			res = mid
			left = mid + 1
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 8\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{4, 2},
		{8, 2},
		{0, 0},
		{1, 1},
		{16, 4},
		{25, 5},
		{100, 10},
	}
	for _, test := range tests {
		if got := mySqrt(test.input); got != test.expected {
			t.Errorf("mySqrt(%d) = %d; want %d", test.input, got, test.expected)
		}
	}
}
