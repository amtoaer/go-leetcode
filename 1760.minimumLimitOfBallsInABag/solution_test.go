package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=1760 lang=golang
 * @lcpr version=30005
 *
 * [1760] 袋子里最少数目的球
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumSize(nums []int, maxOperations int) int {
	var right, res int
	for _, num := range nums {
		right = max(right, num)
	}
	left := 1
	for left <= right {
		mid := left + (right-left)/2
		var operations int
		for _, num := range nums {
			operations += (num - 1) / mid
		}
		if operations <= maxOperations {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [9]\n2\n
// @lcpr case=end

// @lcpr case=start
// [2,4,8,2]\n4\n
// @lcpr case=end

// @lcpr case=start
// [7,17]\n2\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		maxOperations int
		want          int
	}{
		{
			name:          "single bag",
			nums:          []int{9},
			maxOperations: 2,
			want:          3,
		},
		{
			name:          "four bags",
			nums:          []int{2, 4, 8, 2},
			maxOperations: 4,
			want:          2,
		},
		{
			name:          "two bags",
			nums:          []int{7, 17},
			maxOperations: 2,
			want:          7,
		},
		{
			name:          "no operations needed",
			nums:          []int{1, 2, 3},
			maxOperations: 0,
			want:          3,
		},
		{
			name:          "equal bags with spare operations",
			nums:          []int{10, 10, 10},
			maxOperations: 5,
			want:          5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumSize(tt.nums, tt.maxOperations)
			if got != tt.want {
				t.Errorf("minimumSize(%v, %d) = %d, want %d", tt.nums, tt.maxOperations, got, tt.want)
			}
		})
	}
}
