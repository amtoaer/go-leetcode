// Created by amtoaer at 2025/11/21 16:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/edit-distance/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for idx := range dp {
		dp[idx] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	word1 := Deserialize[string](ReadLine(stdin))
	word2 := Deserialize[string](ReadLine(stdin))
	ans := minDistance(word1, word2)

	fmt.Println("\noutput:", Serialize(ans))
}
