package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1
		res += num & 1
		num >>= 1
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  uint32
		output uint32
	}{
		{input: 43261596, output: 964176192},
		{input: 4294967293, output: 3221225471},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, reverseBits(tt.input))
	}
}
