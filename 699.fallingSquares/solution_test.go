package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=699 lang=golang
 *
 * [699] Falling Squares
 */

// @lc code=start
func fallingSquares(positions [][]int) []int {
	height := make([]int, len(positions))
	for idx, pos := range positions {
		left, right := pos[0], pos[0]+pos[1]-1
		height[idx] = pos[1]
		for pervIdx, prevPos := range positions[:idx] {
			prevLeft, prevRight := prevPos[0], prevPos[0]+prevPos[1]-1
			if right < prevLeft || left > prevRight {
				continue
			}
			height[idx] = max(height[idx], height[pervIdx]+pos[1])
		}
	}
	for idx := 1; idx < len(height); idx++ {
		height[idx] = max(height[idx], height[idx-1])
	}
	return height
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]int
		output []int
	}{
		{input: [][]int{{1, 2}, {2, 3}, {6, 1}}, output: []int{2, 5, 5}},
		{input: [][]int{{100, 100}, {200, 100}}, output: []int{100, 100}},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, fallingSquares(tt.input))
	}
}
