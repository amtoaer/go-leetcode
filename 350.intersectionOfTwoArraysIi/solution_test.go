package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, v := range nums1 {
		m[v]++
	}
	var res []int
	for _, v := range nums2 {
		if _, ok := m[v]; ok && m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}
	for _, tt := range tc {
		assert.ElementsMatch(t, tt.want, intersect(tt.nums1, tt.nums2))
	}
}
