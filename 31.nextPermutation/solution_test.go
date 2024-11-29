package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=31 lang=golang
 * @lcpr version=20004
 *
 * [31] 下一个排列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func nextPermutation(nums []int) {
	var left, right int
	for left = len(nums) - 2; left >= 0; left-- {
		if nums[left] < nums[left+1] {
			break
		}
	}
	if left >= 0 {
		for right = len(nums) - 1; right >= 0; right-- {
			if nums[right] > nums[left] {
				break
			}
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	i, j := left+1, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,5]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3}, expected: []int{1, 3, 2}},
		{input: []int{3, 2, 1}, expected: []int{1, 2, 3}},
		{input: []int{1, 1, 5}, expected: []int{1, 5, 1}},
		{input: []int{1, 3, 2}, expected: []int{2, 1, 3}},
		{input: []int{2, 3, 1}, expected: []int{3, 1, 2}},
	}

	for _, test := range tests {
		nextPermutation(test.input)
		for i, v := range test.input {
			if v != test.expected[i] {
				t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, test.input)
				break
			}
		}
	}
}
