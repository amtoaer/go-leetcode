// Created by amtoaer at 2025/11/27 19:36
// leetgo: 1.4.15
// https://leetcode.cn/problems/decode-string/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func decodeString(s string) string {
	var numStack []int
	var letterStack []string
	start := 0
	for start < len(s) {
		if isNumber(s[start]) {
			var num, end int
			for end = start; end < len(s) && isNumber(s[end]); end++ {
				num = num*10 + int(s[end]-'0')
			}
			numStack = append(numStack, num)
			start = end
		} else if isLetter(s[start]) {
			var end int
			for end = start; end < len(s) && isLetter(s[end]); end++ {
			}
			letterStack = append(letterStack, s[start:end])
			start = end
		} else if s[start] == '[' {
			letterStack = append(letterStack, "[")
			start++
		} else {
			// ']'
			var idx int
			for idx = len(letterStack) - 1; idx >= 0 && letterStack[idx] != "["; idx-- {
			}
			strs := letterStack[idx+1:]
			letterStack = letterStack[:idx]
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			var sb strings.Builder
			for range num {
				for _, str := range strs {
					sb.WriteString(str)
				}
			}
			letterStack = append(letterStack, sb.String())
			start++
		}
	}
	if len(letterStack) == 1 {
		return letterStack[0]
	}
	var sb strings.Builder
	for _, str := range letterStack {
		sb.WriteString(str)
	}
	return sb.String()
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := decodeString(s)

	fmt.Println("\noutput:", Serialize(ans))
}
