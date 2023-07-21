package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		var cur strings.Builder
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, "1", countAndSay(1))
	assert.Equal(t, "11", countAndSay(2))
	assert.Equal(t, "21", countAndSay(3))
	assert.Equal(t, "1211", countAndSay(4))
}
