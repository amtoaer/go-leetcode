package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=289 lang=golang
 * @lcpr version=20003
 *
 * [289] 生命游戏
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

var directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func gameOfLife(board [][]int) {
	countAlive := func(i, j int) int {
		var cnt int
		for _, direction := range directions {
			iTarget, jTarget := i+direction[0], j+direction[1]
			if !(iTarget >= 0 && iTarget < len(board) && jTarget >= 0 && jTarget < len(board[0])) {
				continue
			}
			if board[iTarget][jTarget]&1 == 1 {
				cnt++
			}
		}
		return cnt
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			var alive int
			cnt := countAlive(i, j)
			if board[i][j]&1 == 1 {
				// 本来是活的
				alive = 1
				if cnt < 2 || cnt > 3 {
					alive = 0
				}
			} else {
				// 本来是死的
				alive = 0
				if cnt == 3 {
					alive = 1
				}
			}
			board[i][j] |= alive << 1
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] >>= 1
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[1,0]]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		board    [][]int
		expected [][]int
	}{
		{
			board: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			board: [][]int{
				{1, 1},
				{1, 0},
			},
			expected: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		{
			board: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
	}

	for _, test := range tests {
		gameOfLife(test.board)
		for i := range test.board {
			for j := range test.board[i] {
				if test.board[i][j] != test.expected[i][j] {
					t.Errorf("expected %v, but got %v", test.expected, test.board)
				}
			}
		}
	}
}
