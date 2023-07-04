package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var (
		res       []string
		bf        bytes.Buffer
		backtrack func(int, int)
	)
	backtrack = func(left, right int) {
		if left == 0 && right == 0 {
			res = append(res, bf.String())
			return
		}
		if left > 0 {
			bf.WriteByte('(')
			backtrack(left-1, right)
			bf.Truncate(bf.Len() - 1)
		}
		if right > left {
			bf.WriteByte(')')
			backtrack(left, right-1)
			bf.Truncate(bf.Len() - 1)
		}
	}
	backtrack(n, n)
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(
		t,
		[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		generateParenthesis(3),
	)
	assert.Equal(t, []string{"()"}, generateParenthesis(1))
}
