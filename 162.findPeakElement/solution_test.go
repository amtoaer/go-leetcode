package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=162 lang=golang
 * @lcpr version=20004
 *
 * [162] 寻找峰值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		gtLeft := mid <= 0 || nums[mid] > nums[mid-1]
		gtRight := mid >= len(nums)-1 || nums[mid] > nums[mid+1]
		if gtLeft && gtRight {
			return mid
		}
		if gtLeft {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,3,5,6,4]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{1, 3, 2, 1}, 1},
	}

	for _, test := range tests {
		result := findPeakElement(test.nums)
		if result != test.expected {
			t.Errorf("For nums %v, expected %d, but got %d", test.nums, test.expected, result)
		}
	}
}
