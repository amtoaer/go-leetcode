package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=3226 lang=golang
 * @lcpr version=20002
 *
 * [3226] 使两个整数相等的位更改次数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minChanges(n int, k int) int {
	var res int
	for i := 0; i < 32; i++ {
		nLast, kLast := n&1, k&1
		n >>= 1
		k >>= 1
		if nLast == kLast {
			continue
		} else if nLast == 1 && kLast == 0 {
			res++
		} else {
			return -1
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 13\n4\n
// @lcpr case=end

// @lcpr case=start
// 21\n21\n
// @lcpr case=end

// @lcpr case=start
// 14\n13\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		n        int
		k        int
		expected int
	}{
		{13, 4, 2},
		{21, 21, 0},
		{14, 13, -1},
		{0, 0, 0},
		{1, 0, 1},
		{2, 3, -1},
		{7, 0, 3},
		{8, 8, 0},
		{15, 0, 4},
	}

	for _, test := range tests {
		result := minChanges(test.n, test.k)
		if result != test.expected {
			t.Errorf("minChanges(%d, %d) = %d; want %d", test.n, test.k, result, test.expected)
		}
	}
}
