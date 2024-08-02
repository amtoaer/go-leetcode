package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	none := struct{}{}
	m := map[int]struct{}{
		n: none,
	}
	for {
		tmp := 0
		for n > 0 {
			tmp += (n % 10) * (n % 10)
			n /= 10
		}
		if tmp == 1 {
			return true
		}
		if _, ok := m[tmp]; ok {
			return false
		}
		m[tmp] = none
		n = tmp
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  int
		output bool
	}{
		{19, true},
		{2, false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, isHappy(tt.input))
	}
}
