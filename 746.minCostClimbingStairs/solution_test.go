/*
 * @lc app=leetcode.cn id=746 lang=golang
 * @lcpr version=30104
 *
 * [746] 使用最小花费爬楼梯
 */

package main

import (
	"testing"
)

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	var a, b int
	for i := 2; i <= len(cost); i++ {
		a, b = b, min(a+cost[i-2], b+cost[i-1])
	}
	return b
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		cost     []int
		expected int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}

	for _, tt := range tests {
		result := minCostClimbingStairs(tt.cost)
		if result != tt.expected {
			t.Errorf("minCostClimbingStairs(%v) = %v; want %v", tt.cost, result, tt.expected)
		}
	}
}

/*
// @lcpr case=start
// [10,15,20]\n
// @lcpr case=end

// @lcpr case=start
// [1,100,1,1,1,100,1,1,100,1]\n
// @lcpr case=end

*/
