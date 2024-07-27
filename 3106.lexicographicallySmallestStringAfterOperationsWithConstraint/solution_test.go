package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3106 lang=golang
 *
 * [3106] Lexicographically Smallest String After Operations With Constraint
 */

// @lc code=start
func getSmallestString(s string, k int) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		dis := int(bytes[i] - 'a')
		dis = min(dis, 26-dis)
		if k >= dis {
			bytes[i] = 'a'
			k -= dis
		} else {
			bytes[i] -= byte(k)
			break
		}
	}
	return string(bytes)
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		s      string
		k      int
		output string
	}{
		{"zbbz", 3, "aaaz"},
		{"xaxcd", 4, "aawcd"},
		{"lol", 0, "lol"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, getSmallestString(tt.s, tt.k))
	}
}
