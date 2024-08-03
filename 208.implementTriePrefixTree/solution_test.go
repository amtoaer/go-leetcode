package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start

type Trie struct {
	end  bool
	next [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	cur := t
	for i := 0; i < len(word); i++ {
		if next := cur.next[word[i]-'a']; next != nil {
			cur = next
		} else {
			x := &Trie{}
			cur.next[word[i]-'a'] = x
			cur = x
		}
	}
	cur.end = true
}

func (t *Trie) Search(word string) bool {
	cur := t
	for i := 0; i < len(word); i++ {
		if next := cur.next[word[i]-'a']; next == nil {
			return false
		} else {
			cur = next
		}
	}
	return cur.end
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	for i := 0; i < len(prefix); i++ {
		if next := cur.next[prefix[i]-'a']; next == nil {
			return false
		} else {
			cur = next
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

func Test(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple")
	if obj.Search("apple") != true {
		t.Fatal("error")
	}
	if obj.Search("app") != false {
		t.Fatal("error")
	}
	if obj.StartsWith("app") != true {
		t.Fatal("error")
	}
	obj.Insert("app")
	if obj.Search("app") != true {
		t.Fatal("error")
	}
}
