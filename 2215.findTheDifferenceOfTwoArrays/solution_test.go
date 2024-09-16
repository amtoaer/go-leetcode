package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2215 lang=golang
 *
 * [2215] Find the Difference of Two Arrays
 */

// @lc code=start
func findDifference(nums1 []int, nums2 []int) [][]int {
	m1, m2 := make(map[int]struct{}), make(map[int]struct{})
	for _, num := range nums1 {
		m1[num] = struct{}{}
	}
	for _, num := range nums2 {
		m2[num] = struct{}{}
	}
	resL, resR := []int{}, []int{}
	for num := range m1 {
		if _, ok := m2[num]; !ok {
			resL = append(resL, num)
		}
	}
	for num := range m2 {
		if _, ok := m1[num]; !ok {
			resR = append(resR, num)
		}
	}
	return [][]int{resL, resR}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums1 []int
		nums2 []int
		out   [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.out, findDifference(tt.nums1, tt.nums2))
	}
}
