// Created by amtoaer at 2025/11/27 17:47
// leetgo: 1.4.15
// https://leetcode.cn/problems/reverse-words-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseWords(s string) string {
	var wordIdx [][2]int
	var left, right int
	for left < len(s) && right < len(s) {
		for left < len(s) && s[left] == ' ' {
			left++
		}
		for right = left; right < len(s) && s[right] != ' '; right++ {
		}
		if left != right {
			wordIdx = append(wordIdx, [2]int{left, right})
		}
		left = right
	}
	var sb strings.Builder
	for i := len(wordIdx) - 1; i >= 0; i-- {
		v := wordIdx[i]
		sb.WriteString(s[v[0]:v[1]])
		if i > 0 {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := reverseWords(s)

	fmt.Println("\noutput:", Serialize(ans))
}
