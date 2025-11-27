// Created by amtoaer at 2025/11/24 21:09
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-valid-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestValidParentheses(s string) (ans int) {
	var left, right int
	for _, r := range s {
		if r == '(' {
			left++
		}
		if r == ')' {
			right++
		}
		if left == right {
			ans = max(ans, left+right)
		}
		if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		if r == '(' {
			left++
		}
		if r == ')' {
			right++
		}
		if left == right {
			ans = max(ans, left+right)
		}
		if left > right {
			left, right = 0, 0
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := longestValidParentheses(s)

	fmt.Println("\noutput:", Serialize(ans))
}
