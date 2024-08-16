package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=909 lang=golang
 *
 * [909] Snakes and Ladders
 */

// @lc code=start
func snakesAndLadders(board [][]int) int {
	n := len(board)
	getIndex := func(location int) [2]int {
		location -= 1
		row, col := location/n, location%n
		var resultCol int
		if row&1 == 1 {
			resultCol = n - 1 - col
		} else {
			resultCol = col
		}
		return [2]int{n - 1 - row, resultCol}
	}
	getLocation := func(idx [2]int) int {
		i, j := idx[0], idx[1]
		row := n - 1 - i
		cnt := row * n
		if row&1 == 1 {
			cnt += n - j - 1
		} else {
			cnt += j
		}
		return cnt + 1
	}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	// 更好的写法是直接存储 location，这样只需要实现 location -> idx 的转换就够了
	var queue [][2]int
	queue = append(queue, getIndex(1))
	var res int
	for len(queue) > 0 {
		res++
		lenQueue := len(queue)
		for _, start := range queue[:lenQueue] {
			startLoc := getLocation(start)
			for i := 1; i <= 6; i++ {
				if startLoc+i > n*n {
					break
				}
				targetIndex := getIndex(startLoc + i)
				teleport := board[targetIndex[0]][targetIndex[1]]
				if teleport > 0 {
					targetIndex = getIndex(teleport)
				}
				if getLocation(targetIndex) == n*n {
					return res
				}
				if visited[targetIndex[0]][targetIndex[1]] {
					continue
				}
				visited[targetIndex[0]][targetIndex[1]] = true
				queue = append(queue, targetIndex)
			}
		}
		queue = queue[lenQueue:]
	}
	return -1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]int
		output int
	}{
		{[][]int{{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}}, 4},
		// {[][]int{{-1, -1, -1}, {-1, 9, 8}, {-1, 8, 9}}, 1},
		// {[][]int{{-1, -1}, {-1, 3}}, 1},
		// {[][]int{{1, 1, -1}, {1, 1, 1}, {-1, 1, 1}}, -1},
		// {[][]int{{-1, -1, -1, -1, 48, 5, -1}, {12, 29, 13, 9, -1, 2, 32}, {-1, -1, 21, 7, -1, 12, 49}, {42, 37, 21, 40, -1, 22, 12}, {42, -1, 2, -1, -1, -1, 6}, {39, -1, 35, -1, -1, 39, -1}, {-1, 36, -1, -1, -1, -1, 5}}, 3},
		// {[][]int{{-1, 46, -1, 73, 54, 8, 36, -1, -1}, {-1, -1, 10, -1, 15, -1, 15, -1, -1}, {-1, 68, -1, -1, 19, -1, -1, 71, -1}, {-1, -1, -1, -1, 5, -1, 10, -1, 53}, {-1, -1, -1, -1, 4, 50, -1, -1, 18}, {-1, -1, 41, -1, 48, 24, -1, 45, -1}, {43, -1, 18, -1, -1, 57, 55, 29, 20}, {60, -1, -1, -1, -1, 52, -1, 26, -1}, {-1, -1, -1, -1, 65, -1, -1, -1, -1}}, 4},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, snakesAndLadders(tt.input))
	}
}
