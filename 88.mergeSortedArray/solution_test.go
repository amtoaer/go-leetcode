package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start

func merge(nums1 []int, m int, nums2 []int, n int) {
	left, right := m-1, n-1
	current := m + n - 1
	for left >= 0 && right >= 0 {
		if nums1[left] > nums2[right] {
			nums1[current] = nums1[left]
			left--
		} else {
			nums1[current] = nums2[right]
			right--
		}
		current--
	}
	for right >= 0 {
		nums1[current] = nums2[right]
		current -= 1
		right -= 1
	}
}

// @lc code=end

func Test(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
}
