package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				obstacleGrid[i][j] = 1
			} else if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
			} else if j == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  [][]int
		output int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, uniquePathsWithObstacles(tt.input))
	}
}
