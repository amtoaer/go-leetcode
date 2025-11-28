// Created by amtoaer at 2025/11/28 01:20
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-path-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minPathSum(grid [][]int) int {
	var tmp int
	for i := range grid {
		tmp += grid[i][0]
		grid[i][0] = tmp
	}
	tmp = 0
	for i := range grid[0] {
		tmp += grid[0][i]
		grid[0][i] = tmp
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	ans := minPathSum(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
