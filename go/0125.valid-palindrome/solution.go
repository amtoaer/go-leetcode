// Created by amtoaer at 2025/12/17 23:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/valid-palindrome/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for ; left < right && !isValid(s[left]); left++ {
		}
		for ; right > left && !isValid(s[right]); right-- {
		}
		if left < right && !isEqual(s[left], s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isValid(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func toLower(b byte) byte {
	if b >= '0' && b <= '9' {
		return b
	}
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

func isEqual(a, b byte) bool {
	return toLower(a) == toLower(b)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := isPalindrome(s)

	fmt.Println("\noutput:", Serialize(ans))
}
