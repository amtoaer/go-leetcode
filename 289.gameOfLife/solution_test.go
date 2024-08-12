package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start

var direction = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			alive := 0
			for _, d := range direction {
				ii, jj := i+d[0], j+d[1]
				if ii < 0 || ii >= len(board) || jj < 0 || jj >= len(board[0]) {
					continue
				}
				if board[ii][jj]&1 == 1 {
					alive++
				}
			}
			self := board[i][j]
			target := 0
			if self == 1 {
				if alive < 2 || alive > 3 {
					target = 0
				} else {
					target = 1
				}
			} else {
				if alive == 3 {
					target = 1
				}
			}
			board[i][j] |= target << 1
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] >>= 1
		}
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}},
			output: [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}},
		},
	}
	for _, tt := range tc {
		gameOfLife(tt.input)
		assert.Equal(t, tt.output, tt.input)
	}
}
