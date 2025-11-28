// Created by amtoaer at 2025/11/28 11:50
// leetgo: 1.4.15
// https://leetcode.cn/problems/max-area-of-island/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxAreaOfIsland(grid [][]int) (ans int) {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i-1, j) + dfs(i, j-1) + dfs(i+1, j) + dfs(i, j+1)
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	ans := maxAreaOfIsland(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
