package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=88 lang=golang
 * @lcpr version=20004
 *
 * [88] 合并两个有序数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	cur := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[cur] = nums1[m]
			m--
		} else {
			nums1[cur] = nums2[n]
			n--
		}
		cur--
	}
	for n >= 0 {
		nums1[cur] = nums2[n]
		cur--
		n--
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,0,0]\n3\n[2,5,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n[]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n[1]\n1\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
	}

	for _, tt := range tests {
		merge(tt.nums1, tt.m, tt.nums2, tt.n)
		for i := range tt.nums1 {
			if tt.nums1[i] != tt.want[i] {
				t.Errorf("got %v, want %v", tt.nums1, tt.want)
				break
			}
		}
	}
}
