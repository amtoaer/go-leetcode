package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=419 lang=golang
 *
 * [419] Battleships in a Board
 */

// @lc code=start
func countBattleships(board [][]byte) int {
	is_corner := func(i, j int) bool {
		return (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.')
	}
	var res int
	x, y := len(board), len(board[0])
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if board[i][j] == 'X' && is_corner(i, j) {
				res++
			}
		}
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		board  [][]byte
		output int
	}{
		{[][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}, 2},
		{[][]byte{{'X', 'X', 'X', 'X'}, {'.', '.', '.', '.'}, {'.', '.', '.', '.'}}, 1},
		{[][]byte{{'.'}}, 0},
	}
	for _, testcase := range testcases {
		assert.Equal(t, testcase.output, countBattleships(testcase.board))
	}
}
