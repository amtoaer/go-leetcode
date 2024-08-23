package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] Remove K Digits
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	var stack []byte
	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > num[i] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	var zeroPos int
	for zeroPos < len(stack) && stack[zeroPos] == '0' {
		zeroPos++
	}
	if zeroPos >= len(stack)-k {
		return "0"
	}
	return string(stack[zeroPos : len(stack)-k])
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  string
		k      int
		output string
	}{
		{"1432219", 3, "1219"},
		{"10200", 1, "200"},
		{"10", 2, "0"},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, removeKdigits(tt.input, tt.k))
	}
}
