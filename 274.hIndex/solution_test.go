package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	sort.Ints(citations)
	for idx, num := range citations {
		greater := len(citations) - idx
		if greater <= num {
			return greater
		}
	}
	return 0
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, hIndex(tt.input))
	}
}
