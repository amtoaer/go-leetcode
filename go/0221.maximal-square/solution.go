// Created by amtoaer at 2025/11/28 14:13
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximal-square/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	var ans int
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans * ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]byte](ReadLine(stdin))
	ans := maximalSquare(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}
