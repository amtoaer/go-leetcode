package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2982 lang=golang
 *
 * [2982] Find Longest Special Substring That Occurs Thrice II
 */

// @lc code=start
func orderedInsert(l []int, v int) []int {
	var i int
	for i < len(l) {
		if l[i] < v {
			break
		}
		i++
	}
	l = append(l, 0)
	copy(l[i+1:], l[i:])
	l[i] = v
	if len(l) > 3 {
		l = l[:3]
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumLength(s string) int {
	var (
		tmp   [26][]int
		count = 0
		res   = -1
	)

	for i := 0; i < len(s); i++ {
		count++
		if i == len(s)-1 || s[i] != s[i+1] {
			tmp[s[i]-'a'] = orderedInsert(tmp[s[i]-'a'], count)
			count = 0
		}
	}

	for i := 0; i < 26; i++ {
		if len(tmp[i]) > 0 && tmp[i][0] > 2 {
			res = max(res, tmp[i][0]-2)
		}
		if len(tmp[i]) > 1 {
			if tmp[i][0] == tmp[i][1] && tmp[i][0] > 1 {
				res = max(res, tmp[i][0]-1)
			} else if tmp[i][0] != tmp[i][1] {
				res = max(res, tmp[i][1])
			}
		}
		if len(tmp[i]) > 2 {
			res = max(res, tmp[i][2])
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	var testcases = []struct {
		s   string
		ans int
	}{
		{"aaaa", 2},
		{"abcdef", -1},
		{"abcaba", 1},
		{"cbc", -1},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.ans, maximumLength(tc.s))
	}
}
