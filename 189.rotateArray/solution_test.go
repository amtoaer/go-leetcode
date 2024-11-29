package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=189 lang=golang
 * @lcpr version=20004
 *
 * [189] 轮转数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rotate(nums []int, k int) {
	reverse := func(i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	k %= len(nums)
	reverse(0, len(nums)-1)
	reverse(0, k-1)
	reverse(k, len(nums)-1)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [-1,-100,3,99]\n2\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
		{[]int{1, 2}, 3, []int{2, 1}},
		{[]int{1}, 0, []int{1}},
	}

	for _, test := range tests {
		rotate(test.nums, test.k)
		for i, v := range test.expected {
			if test.nums[i] != v {
				t.Errorf("For input nums=%v and k=%d, expected %v but got %v", test.nums, test.k, test.expected, test.nums)
				break
			}
		}
	}
}
