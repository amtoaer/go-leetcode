// Created by amtoaer at 2025/11/18 15:07
// leetgo: 1.4.15
// https://leetcode.cn/problems/valid-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isValid(s string) bool {
	var stack []rune
	pair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			{
				stack = append(stack, r)
			}
		default:
			{
				target := pair[r]
				if len(stack) == 0 || stack[len(stack)-1] != target {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := isValid(s)

	fmt.Println("\noutput:", Serialize(ans))
}
