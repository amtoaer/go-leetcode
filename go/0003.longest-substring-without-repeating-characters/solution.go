// Created by amtoaer at 2025/11/17 21:33
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLongestSubstring(s string) (ans int) {
	m := make(map[byte]any)
	l := 0
	for r, b := range []byte(s) {
		for {
			if _, ok := m[b]; !ok {
				break
			}
			delete(m, s[l])
			l++
		}
		m[b] = struct{}{}
		ans = max(ans, r-l+1)
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLongestSubstring(s)

	fmt.Println("\noutput:", Serialize(ans))
}
