package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start
func lexicalOrder(n int) []int {
	var res []int
	var dfs func(int)
	dfs = func(base int) {
		for i := 0; i <= 9; i++ {
			v := base + i
			if v == 0 {
				continue
			}
			if v > n {
				return
			}
			res = append(res, v)
			dfs(v * 10)
		}
	}
	dfs(0)
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  int
		output []int
	}{
		{13, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		{1, []int{1}},
		{2, []int{1, 2}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, lexicalOrder(tt.input))
	}
}
