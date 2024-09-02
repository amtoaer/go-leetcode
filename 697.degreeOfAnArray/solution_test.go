package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] Degree of an Array
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	m := make(map[int][]int)
	for idx, num := range nums {
		if v, ok := m[num]; !ok {
			m[num] = []int{1, idx, idx}
		} else {
			v[0]++
			v[2] = idx
		}
	}
	var maxV int
	for _, r := range m {
		maxV = max(maxV, r[0])
	}
	res := math.MaxInt
	for _, r := range m {
		if r[0] == maxV {
			res = min(res, r[2]-r[1]+1)
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
		{[]int{1}, 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, findShortestSubArray(tt.nums))
	}
}
