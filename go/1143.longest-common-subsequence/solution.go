// Created by amtoaer at 2025/11/21 17:08
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-common-subsequence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for idx := range dp {
		dp[idx] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	text1 := Deserialize[string](ReadLine(stdin))
	text2 := Deserialize[string](ReadLine(stdin))
	ans := longestCommonSubsequence(text1, text2)

	fmt.Println("\noutput:", Serialize(ans))
}
