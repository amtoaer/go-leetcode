package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=215 lang=golang
 * @lcpr version=20003
 *
 * [215] 数组中的第K个最大元素
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func siftDown(nums []int, start, end int) {
	cur, next := start, start*2+1
	for next <= end {
		if next+1 <= end && nums[next+1] > nums[next] {
			next = next + 1
		}
		if nums[cur] >= nums[next] {
			break
		}
		nums[cur], nums[next] = nums[next], nums[cur]
		cur, next = next, next*2+1
	}
}

func findKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, i, len(nums)-1)
	}
	for i := 0; i < k; i++ {
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		siftDown(nums, 0, len(nums)-i-2)
	}
	return nums[len(nums)-k]
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{1}, 1, 1},
		{[]int{2, 1}, 1, 2},
		{[]int{2, 1}, 2, 1},
	}

	for _, test := range tests {
		result := findKthLargest(test.nums, test.k)
		if result != test.expected {
			t.Errorf("findKthLargest(%v, %d) = %d; expected %d", test.nums, test.k, result, test.expected)
		}
	}
}
