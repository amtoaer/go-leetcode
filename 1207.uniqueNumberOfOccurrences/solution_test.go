package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] Unique Number of Occurrences
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	cnt := make(map[int]int)
	for _, v := range arr {
		cnt[v]++
	}
	m := make(map[int]struct{})
	for _, v := range cnt {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, uniqueOccurrences(tt.input))
	}
}
