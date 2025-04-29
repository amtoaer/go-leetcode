/*
 * @lc app=leetcode.cn id=2962 lang=golang
 * @lcpr version=30200
 *
 * [2962] 统计最大元素出现至少 K 次的子数组
 */

package main

import (
	"math"
	"testing"
)

// @lc code=start
func countSubarrays(nums []int, k int) int64 {
	maxVal := math.MinInt32
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	var (
		left, cnt int
		res       int64
	)
	for _, rightVal := range nums {
		if rightVal == maxVal {
			cnt++
		}
		for cnt == k {
			if nums[left] == maxVal {
				cnt--
			}
			left++
		}
		res += int64(left)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int64
	}{
		{nums: []int{1, 3, 2, 3, 3}, k: 2, want: 6},
		{nums: []int{1, 4, 2, 1}, k: 3, want: 0},
	}

	for _, tt := range tests {
		got := countSubarrays(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("countSubarrays(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}

/*
// @lcpr case=start
// [1,3,2,3,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,4,2,1]\n3\n
// @lcpr case=end

*/
