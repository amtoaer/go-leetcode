// Created by amtoaer at 2025/11/18 02:20
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-palindromic-substring/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	bytes := []byte(s)
	var ans []byte
	for i := 0; i < len(bytes)-1; i++ {
		for _, r := range [2]int{0, 1} {
			tmpAns := palindrome(bytes, i, i+r)
			if len(tmpAns) > len(ans) {
				ans = tmpAns
			}
		}
	}
	return string(ans)
}

func palindrome(s []byte, left, right int) []byte {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	if left+1 > right {
		return nil
	}
	return s[left+1 : right]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := longestPalindrome(s)

	fmt.Println("\noutput:", Serialize(ans))
}
