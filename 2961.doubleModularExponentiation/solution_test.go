package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2961 lang=golang
 *
 * [2961] Double Modular Exponentiation
 */

// @lc code=start
func getGoodIndices(variables [][]int, target int) []int {
	res := []int{}
	for i, v := range variables {
		if quickMod(quickMod(v[0], v[1], 10), v[2], v[3]) == target {
			res = append(res, i)
		}
	}
	return res
}

func quickMod(a, b, mod int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		variables [][]int
		target    int
		expect    []int
	}{
		{
			variables: [][]int{{2, 3, 3, 10}, {3, 3, 3, 1}, {6, 1, 1, 4}},
			target:    2,
			expect:    []int{0, 2},
		},
		{
			variables: [][]int{{39, 3, 1000, 1000}},
			target:    17,
			expect:    []int{},
		},
		// 溢出的情况
		{
			variables: [][]int{
				{30, 5, 43, 2},
				{15, 50, 35, 41},
				{45, 34, 41, 32},
				{14, 37, 33, 13},
				{6, 8, 1, 53},
				{37, 1, 12, 52},
				{42, 37, 2, 52},
				{9, 2, 15, 3},
				{31, 12, 21, 24},
				{52, 24, 6, 12},
				{51, 35, 21, 52},
				{30, 18, 10, 2},
				{27, 31, 50, 27},
				{29, 25, 26, 32},
				{15, 38, 43, 17},
				{22, 12, 16, 43},
				{48, 9, 15, 6},
				{41, 26, 22, 21},
				{41, 49, 52, 26},
				{53, 38, 9, 33},
			},
			target: 1,
			expect: []int{5, 7, 8, 10, 17, 18},
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.expect, getGoodIndices(tt.variables, tt.target))
	}
}
