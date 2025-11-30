// Created by amtoaer at 2025/11/30 16:04
// leetgo: 1.4.15
// https://leetcode.cn/problems/word-break/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]any)
	for _, word := range wordDict {
		m[word] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := range i {
			if _, ok := m[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	wordDict := Deserialize[[]string](ReadLine(stdin))
	ans := wordBreak(s, wordDict)

	fmt.Println("\noutput:", Serialize(ans))
}
