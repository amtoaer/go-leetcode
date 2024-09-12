package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=1431 lang=golang
 *
 * [1431] Kids With the Greatest Number of Candies
 */

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxV := math.MinInt
	for _, candie := range candies {
		maxV = max(maxV, candie)
	}
	res := make([]bool, len(candies))
	for idx, candie := range candies {
		res[idx] = candie+extraCandies >= maxV
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		candies      []int
		extraCandies int
		expect       []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.expect, kidsWithCandies(tt.candies, tt.extraCandies))
	}
}
