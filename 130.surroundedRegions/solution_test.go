package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
var directions = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func solve(board [][]byte) {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'A'
		for _, direction := range directions {
			dfs(i+direction[0], j+direction[1])
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			dfs(i, len(board[0])-1)
		}
	}
	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[len(board)-1][j] == 'O' {
			dfs(len(board)-1, j)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]byte
		output [][]byte
	}{
		{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}, [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		{[][]byte{{'X'}}, [][]byte{{'X'}}},
		{[][]byte{{'O'}}, [][]byte{{'O'}}},
	}
	for _, tt := range tc {
		solve(tt.input)
		assert.True(t, reflect.DeepEqual(tt.input, tt.output))
	}
}
