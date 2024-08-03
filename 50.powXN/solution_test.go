package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	var label bool
	if n < 0 {
		n = -n
		label = true
	}
	res := 1.0
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	if label {
		return 1 / res
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	almostEqual := func(a, b float64) bool {
		return (a-b) < 1e-9 && (b-a) < 1e-9
	}

	tc := []struct {
		input1 float64
		input2 int
		output float64
	}{
		{input1: 2.00000, input2: 10, output: 1024.00000},
		{input1: 2.10000, input2: 3, output: 9.26100},
		{input1: 2.00000, input2: -2, output: 0.25000},
	}
	for _, tt := range tc {
		assert.True(t, almostEqual(tt.output, myPow(tt.input1, tt.input2)))
	}
}
