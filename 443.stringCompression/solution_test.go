package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start
func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	var left, right int
	for right < len(chars) {
		char, count := chars[right], 0
		for right < len(chars) && chars[right] == char {
			right++
			count++
		}
		chars[left] = char
		left++
		if count > 1 {
			for _, c := range []byte(strconv.Itoa(count)) {
				chars[left] = c
				left++
			}
		}
	}
	return left
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []byte
		output int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{[]byte{'a'}, 1},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, compress(tt.input))
	}
}
