// Created by amtoaer at 2025/11/26 15:34
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-window-substring/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minWindow(s string, t string) string {
	ms, mt := make(map[byte]int), make(map[byte]int)
	for i := range t {
		mt[t[i]]++
	}
	var left, count int
	bytesS := []byte(s)
	var ans []byte
	for right := range bytesS {
		b := s[right]
		ms[b]++
		if v, ok := mt[b]; ok && v == ms[b] {
			count++
		}
		for count == len(mt) {
			if len(ans) == 0 || right-left+1 < len(ans) {
				ans = bytesS[left : right+1]
			}
			lb := s[left]
			left++
			if v, ok := mt[lb]; ok && v == ms[lb] {
				count--
			}
			ms[lb]--
		}
	}
	return string(ans)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := minWindow(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
