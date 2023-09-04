package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2605 lang=golang
 *
 * [2605] Form Smallest Number From Two Digit Arrays
 */

// @lc code=start
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minNumber(nums1 []int, nums2 []int) int {
	m := make(map[int]struct{})
	res, left, right := math.MaxInt, math.MaxInt, math.MaxInt
	for _, num := range nums1 {
		left = min(left, num)
		m[num] = struct{}{}
	}
	for _, num := range nums2 {
		right = min(right, num)
		if _, ok := m[num]; ok {
			res = min(res, num)
		}
	}
	if left > right {
		left, right = right, left
	}
	return min(res, left*10+right)
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 15, minNumber([]int{4, 1, 3}, []int{5, 7}))
	assert.Equal(t, 3, minNumber([]int{3, 5, 2, 6}, []int{3, 1, 7}))
}
