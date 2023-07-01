package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	var sb strings.Builder
	space := 2*numRows - 2
	if space == 0 {
		return s
	}
	for i := 0; i < numRows; i++ {
		j := i
		for j < len(s) {
			if i == 0 || i == numRows-1 {
				sb.WriteByte(s[j])
				j += space
			} else {
				sb.WriteByte(s[j])
				j += space - 2*i
				if j < len(s) {
					sb.WriteByte(s[j])
					j += 2 * i
				}
			}
		}
	}
	return sb.String()
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR")
	assert.Equal(t, convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI")
}
