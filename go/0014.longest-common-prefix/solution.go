// Created by amtoaer at 2025/11/29 00:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-common-prefix/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type Trie struct {
	next [26]*Trie
}

func InitTrie(s string) *Trie {
	root := &Trie{}
	cur := root
	for _, r := range s {
		idx := r - 'a'
		if cur.next[idx] == nil {
			cur.next[idx] = &Trie{}
		}
		cur = cur.next[idx]
	}
	return root
}

func (t *Trie) Lookup(s string) (ans int) {
	cur := t
	for _, r := range s {
		idx := r - 'a'
		if cur.next[idx] == nil {
			return ans
		}
		cur = cur.next[idx]
		ans++
	}
	return
}

func longestCommonPrefix(strs []string) string {
	var s string
	minLen := math.MaxInt64
	for _, str := range strs {
		if len(str) < minLen {
			s = str
			minLen = len(str)
		}
	}
	if minLen == 0 {
		return ""
	}
	sTrie := InitTrie(s)
	target := len(s)
	for _, str := range strs {
		target = min(target, sTrie.Lookup(str))
	}
	return s[:target]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	strs := Deserialize[[]string](ReadLine(stdin))
	ans := longestCommonPrefix(strs)

	fmt.Println("\noutput:", Serialize(ans))
}
