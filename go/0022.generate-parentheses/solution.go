// Created by amtoaer at 2025/11/22 11:21
// leetgo: 1.4.15
// https://leetcode.cn/problems/generate-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func generateParenthesis(n int) (ans []string) {
	var bytes []byte
	var backtrack func(int, int)
	backtrack = func(left, right int) {
		if left == n && right == n {
			ans = append(ans, string(bytes))
			return
		}
		if left <= n {
			bytes = append(bytes, '(')
			backtrack(left+1, right)
			bytes = bytes[:len(bytes)-1]
		}
		if right < left {
			bytes = append(bytes, ')')
			backtrack(left, right+1)
			bytes = bytes[:len(bytes)-1]
		}
	}
	backtrack(0, 0)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := generateParenthesis(n)

	fmt.Println("\noutput:", Serialize(ans))
}
