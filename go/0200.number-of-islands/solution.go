// Created by amtoaer at 2025/11/19 22:42
// leetgo: 1.4.15
// https://leetcode.cn/problems/number-of-islands/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numIslands(grid [][]byte) (ans int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}
	var visit func(i, j int)
	visit = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0'
		for _, direction := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			visit(i+direction[0], j+direction[1])
		}
	}
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				ans++
				visit(i, j)
			}
		}
	}
	return
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]byte](ReadLine(stdin))
	ans := numIslands(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
