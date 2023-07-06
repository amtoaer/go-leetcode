package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2178 lang=golang
 *
 * [2178] Maximum Split of Positive Even Integers
 */

// @lc code=start
func maximumEvenSplit(finalSum int64) []int64 {
	res := []int64{}
	if finalSum&1 == 1 {
		return res
	}
	target := finalSum
	for i := int64(2); i <= target; i += 2 {
		res = append(res, i)
		target -= i
	}
	res[len(res)-1] += target
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, []int64{2, 4, 6}, maximumEvenSplit(12))
	assert.Equal(t, []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, maximumEvenSplit(110))
	assert.Equal(t, []int64{}, maximumEvenSplit(7))
	assert.Equal(t, []int64{2, 4, 6, 16}, maximumEvenSplit(28))
}
