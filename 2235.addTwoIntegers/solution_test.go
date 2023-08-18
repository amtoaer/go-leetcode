package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2235 lang=golang
 *
 * [2235] Add Two Integers
 */

// @lc code=start
func sum(num1 int, num2 int) int {
	return num1 + num2
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 3, sum(1, 2))
	assert.Equal(t, 5, sum(2, 3))
}
