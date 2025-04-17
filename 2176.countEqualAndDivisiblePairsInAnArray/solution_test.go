/*
 * @lc app=leetcode.cn id=2176 lang=golang
 * @lcpr version=30104
 *
 * [2176] 统计数组中相等且可以被整除的数对
 */

package main

import (
	"testing"
)

// @lc code=start
func countPairs(nums []int, k int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i*j%k == 0 {
				res++
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 1, 2, 2, 2, 1, 3}, 2, 4},
		{[]int{1, 2, 3, 4}, 1, 0},
	}

	for _, tt := range tests {
		got := countPairs(tt.nums, tt.k)
		if got != tt.expected {
			t.Errorf("countPairs(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.expected)
		}
	}
}

/*
// @lcpr case=start
// [3,1,2,2,2,1,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n1\n
// @lcpr case=end

*/
