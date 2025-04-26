/*
 * @lc app=leetcode.cn id=416 lang=golang
 * @lcpr version=30200
 *
 * [416] 分割等和子集
 */

package main

import (
	"testing"
)

// @lc code=start
func canPartition(nums []int) bool {
	var sum, maxVal int
	for _, num := range nums {
		sum += num
		maxVal = max(num, maxVal)
	}
	target := sum / 2
	if sum%2 != 0 || maxVal > target {
		return false
	}
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[target]
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
	}
	for _, test := range tests {
		result := canPartition(test.nums)
		if result != test.expected {
			t.Errorf("canPartition(%v) = %v; want %v", test.nums, result, test.expected)
		}
	}
}

/*
// @lcpr case=start
// [1,5,11,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,5]\n
// @lcpr case=end

*/
