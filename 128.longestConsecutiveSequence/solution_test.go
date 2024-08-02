package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	none := struct{}{}
	m := make(map[int]struct{})
	for _, num := range nums {
		m[num] = none
	}
	res := 0
	for k := range m {
		if _, ok := m[k-1]; ok {
			continue
		}
		cur := k
		cnt := 1
		for {
			if _, ok := m[cur+1]; !ok {
				break
			}
			cur++
			cnt++
		}
		res = max(res, cnt)
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, longestConsecutive(tt.input))
	}
}
