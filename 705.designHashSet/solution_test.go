package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] Design HashSet
 */

// @lc code=start
type Node struct {
	key  int
	next *Node
}

type MyHashSet struct {
	hashSet []*Node
}

func Constructor() MyHashSet {
	return MyHashSet{
		hashSet: make([]*Node, 1009),
	}
}

func (m *MyHashSet) hash(key int) int {
	return key % len(m.hashSet)
}

func (m *MyHashSet) Add(key int) {
	idx := m.hash(key)
	if m.hashSet[idx] == nil {
		m.hashSet[idx] = &Node{
			key: -1,
		}
	}
	prev := m.hashSet[idx]
	for cur := prev.next; cur != nil; prev, cur = cur, cur.next {
		if cur.key == key {
			return
		}
	}
	prev.next = &Node{
		key: key,
	}
}

func (m *MyHashSet) Remove(key int) {
	idx := m.hash(key)
	var prev *Node
	for cur := m.hashSet[idx]; cur != nil; prev, cur = cur, cur.next {
		if cur.key == key {
			prev.next = cur.next
			return
		}
	}
}

func (m *MyHashSet) Contains(key int) bool {
	idx := m.hash(key)
	for cur := m.hashSet[idx]; cur != nil; cur = cur.next {
		if cur.key == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

func Test(t *testing.T) {
}
