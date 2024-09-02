package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=3127 lang=golang
 *
 * [3127] Make a Square with the Same Color
 */

// @lc code=start

func canMakeSquare(grid [][]byte) bool {
	l := len(grid)
	check := func(i, j int) bool {
		if i+1 >= l || j+1 >= l {
			return false
		}
		var cnt [2]int
		for x := i; x < i+2; x++ {
			for y := j; y < j+2; y++ {
				if grid[x][y] == 'B' {
					cnt[0] += 1
				} else {
					cnt[1] += 1
				}
			}
		}
		if cnt[0] == 2 || cnt[1] == 2 {
			return false
		}
		return true
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if check(i, j) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]byte
		expect bool
	}{
		{
			input: [][]byte{
				{'B', 'W', 'B'},
				{'B', 'W', 'W'},
				{'B', 'W', 'B'},
			},
			expect: true,
		},
		{
			input: [][]byte{
				{'B', 'W', 'B'},
				{'W', 'B', 'W'},
				{'B', 'W', 'B'},
			},
			expect: false,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, canMakeSquare(tt.input), tt.expect)
	}
}
