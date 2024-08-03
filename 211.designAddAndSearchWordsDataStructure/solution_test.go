package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */

// @lc code=start
type WordDictionary struct {
	end  bool
	next [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (w *WordDictionary) AddWord(word string) {
	cur := w
	for i := 0; i < len(word); i++ {
		if next := cur.next[word[i]-'a']; next != nil {
			cur = next
		} else {
			tmp := &WordDictionary{}
			cur.next[word[i]-'a'] = tmp
			cur = tmp
		}
	}
	cur.end = true
}

func (w *WordDictionary) Search(word string) bool {
	cur := []*WordDictionary{w}
	for i := 0; i < len(word); i++ {
		curLen := len(cur)
		if curLen == 0 {
			return false
		}
		if word[i] != '.' {
			idx := word[i] - 'a'
			for _, item := range cur[:curLen] {
				if next := item.next[idx]; next != nil {
					cur = append(cur, next)
				}
			}
		} else {
			for _, item := range cur[:curLen] {
				for _, next := range item.next {
					if next != nil {
						cur = append(cur, next)
					}
				}
			}
		}
		cur = cur[curLen:]
	}
	for _, item := range cur {
		if item.end {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

func Test(t *testing.T) {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	assert.False(t, obj.Search("pad"))
	assert.True(t, obj.Search("bad"))
	assert.True(t, obj.Search(".ad"))
	assert.True(t, obj.Search("b.."))
}
