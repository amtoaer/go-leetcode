package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=70 lang=golang
 * @lcpr version=20004
 *
 * [70] 爬楼梯
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	first, second := 1, 2
	for i := 2; i < n; i++ {
		first, second = second, first+second
	}
	return second
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	}

	for _, test := range tests {
		if result := climbStairs(test.n); result != test.expected {
			t.Errorf("climbStairs(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}
