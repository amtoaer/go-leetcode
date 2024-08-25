package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var lp, rp int
	for lp < len(g) && rp < len(s) {
		if s[rp] >= g[lp] {
			lp++
		}
		rp++
	}
	return lp
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		g  []int
		s  []int
		ex int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.ex, findContentChildren(tt.g, tt.s))
	}
}
