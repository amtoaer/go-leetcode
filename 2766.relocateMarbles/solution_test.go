package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2766 lang=golang
 *
 * [2766] Relocate Marbles
 */

// @lc code=start
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	d := make(map[int]struct{})
	for _, n := range nums {
		d[n] = struct{}{}
	}
	for i := 0; i < len(moveFrom); i++ {
		if _, ok := d[moveFrom[i]]; ok {
			delete(d, moveFrom[i])
			d[moveTo[i]] = struct{}{}
		}
	}
	keys := make([]int, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		nums     []int
		moveFrom []int
		moveTo   []int
		want     []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2}, []int{3, 4}, []int{3, 4, 5}},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, relocateMarbles(tc.nums, tc.moveFrom, tc.moveTo))
	}
}
