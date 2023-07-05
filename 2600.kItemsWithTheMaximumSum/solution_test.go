package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2600 lang=golang
 *
 * [2600] K Items With the Maximum Sum
 */

// @lc code=start
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type group struct {
	count int
	value int
}

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	all := []group{
		{numOnes, 1},
		{numZeros, 0},
		{numNegOnes, -1},
	}
	var sum int
	for _, item := range all {
		used := min(k, item.count)
		sum += used * item.value
		k -= used
		if k <= 0 {
			break
		}
	}
	return sum
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, 1, kItemsWithMaximumSum(1, 0, 0, 1))
	assert.Equal(t, 1, kItemsWithMaximumSum(1, 1, 0, 2))
	assert.Equal(t, 0, kItemsWithMaximumSum(1, 1, 1, 3))
}
