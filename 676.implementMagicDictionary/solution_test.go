package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=676 lang=golang
 *
 * [676] Implement Magic Dictionary
 */

// @lc code=start
type MagicDictionary struct {
	dict map[byte]*MagicDictionary
	end  bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		dict: make(map[byte]*MagicDictionary),
	}
}

func (m *MagicDictionary) insertWord(word string) {
	cur := m
	for i := 0; i < len(word); i++ {
		dict, ok := cur.dict[word[i]]
		if !ok {
			child := Constructor()
			dict = &child
			cur.dict[word[i]] = dict
		}
		cur = dict
	}
	cur.end = true
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		m.insertWord(word)
	}
}

func (m *MagicDictionary) Search(searchWord string) bool {
	var backtrack func(cur *MagicDictionary, word string, modified bool) bool
	backtrack = func(cur *MagicDictionary, word string, modified bool) bool {
		if len(word) == 0 {
			return modified && cur.end
		}
		next, ok := cur.dict[word[0]]
		if ok && backtrack(next, word[1:], modified) {
			return true
		}
		if !modified {
			for _, d := range cur.dict {
				if d != next && backtrack(d, word[1:], true) {
					return true
				}
			}
		}
		return false
	}
	return backtrack(m, searchWord, false)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
// @lc code=end

func Test(t *testing.T) {

	// ["MagicDictionary", "buildDict", "search", "search", "search", "search"]
	// [[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
	// 输出
	// [null, null, false, true, false, false]

	m := Constructor()
	m.BuildDict([]string{"hello", "leetcode"})
	assert.False(t, m.Search("hello"))
	assert.True(t, m.Search("hhllo"))
	assert.False(t, m.Search("hell"))
	assert.False(t, m.Search("leetcoded"))

}
