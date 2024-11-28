package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=26 lang=golang
 * @lcpr version=20003
 *
 * [26] 删除有序数组中的重复项
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func removeDuplicates(nums []int) int {
	var left int
	for right := 0; right < len(nums); right++ {
		if right == 0 || nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,2,2,3,3,4]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{1, 1, 1, 1, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := removeDuplicates(test.input)
		if result != test.expected {
			t.Errorf("removeDuplicates(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
