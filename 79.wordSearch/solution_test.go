package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	var check func(int, int, string) bool
	check = func(i, j int, match string) bool {
		if len(match) == 0 {
			return true
		}
		if !(i >= 0 && i < len(board) && j >= 0 && j < len(board[0])) ||
			(board[i][j] == '0') ||
			(board[i][j] != match[0]) {
			return false
		}
		board[i][j] = '0'
		defer func() { board[i][j] = match[0] }()
		for _, diff := range [][]int{
			{0, 1},
			{0, -1},
			{1, 0},
			{-1, 0},
		} {
			if check(i+diff[0], j+diff[1], match[1:]) {
				return true
			}
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if check(i, j, word) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCCED", true},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "SEE", true},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCB", false},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, exist(tt.board, tt.word))
	}
}
