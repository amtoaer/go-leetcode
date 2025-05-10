/*
 * @lc app=leetcode.cn id=1137 lang=golang
 * @lcpr version=30200
 *
 * [1137] 第 N 个泰波那契数
 */

package main

import (
	"testing"
)

// @lc code=start
func tribonacci(n int) int {
	var nums = [3]int{0, 1, 1}
	if n <= 2 {
		return nums[n]
	}
	for range n - 2 {
		nums[0], nums[1], nums[2] = nums[1], nums[2], nums[0]+nums[1]+nums[2]
	}
	return nums[2]
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{n: 4, expected: 4},
		{n: 25, expected: 1389537},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tribonacci(tt.n); got != tt.expected {
				t.Errorf("tribonacci(%d) = %d, want %d", tt.n, got, tt.expected)
			}
		})
	}
}

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 25\n
// @lcpr case=end

*/
