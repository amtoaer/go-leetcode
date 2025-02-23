package main

import (
	"math/rand"
	"testing"
)

/*
 * @lc app=leetcode.cn id=1206 lang=golang
 * @lcpr version=30006
 *
 * [1206] 设计跳表
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

const maxLevel = 32
const pFactor = 0.5

type SkiplistNode struct {
	Val     int
	forward []*SkiplistNode
}

type Skiplist struct {
	level int
	head  *SkiplistNode
}

func Constructor() Skiplist {
	return Skiplist{level: 0, head: &SkiplistNode{Val: -1, forward: make([]*SkiplistNode, maxLevel)}}
}

func (Skiplist) randomLevel() int {
	level := 1
	for level < maxLevel && rand.Float64() < pFactor {
		level++
	}
	return level
}

func (s *Skiplist) Search(target int) bool {
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].Val < target {
			cur = cur.forward[i]
		}
	}
	next := cur.forward[0]
	if next != nil && next.Val == target {
		return true
	}
	return false
}

func (s *Skiplist) Add(num int) {
	update := make([]*SkiplistNode, maxLevel)
	for i := range update {
		update[i] = s.head
	}
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].Val < num {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	level := s.randomLevel()
	s.level = max(s.level, level)
	newNode := &SkiplistNode{Val: num, forward: make([]*SkiplistNode, level)}
	for idx, node := range update[:level] {
		newNode.forward[idx] = node.forward[idx]
		node.forward[idx] = newNode
	}
}

func (s *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, maxLevel)
	for i := range update {
		update[i] = s.head
	}
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].Val < num {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	next := cur.forward[0]
	if next == nil || next.Val != num {
		return false
	}
	for i := 0; i < s.level && update[i].forward[i] == next; i++ {
		update[i].forward[i] = next.forward[i]
	}
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}
	return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
// @lc code=end

func Test(t *testing.T) {
	// Test 1: Search in empty Skiplist.
	sl := Constructor()
	if sl.Search(5) {
		t.Error("Search in an empty skiplist should return false")
	}

	// Test 2: Add one element and search.
	sl.Add(5)
	if !sl.Search(5) {
		t.Error("Search should return true after adding 5")
	}

	// Test 3: Add multiple elements and verify search.
	numsToAdd := []int{1, 3, 7, 9, 12}
	for _, num := range numsToAdd {
		sl.Add(num)
	}
	for _, num := range numsToAdd {
		if !sl.Search(num) {
			t.Errorf("Search should return true for %d", num)
		}
	}

	// Test 4: Erase an element and confirm that it is removed.
	if !sl.Erase(3) {
		t.Error("Erase should return true for existing element 3")
	}
	if sl.Search(3) {
		t.Error("Search should return false after erasing 3")
	}

	// Test 5: Attempt to erase an element that does not exist.
	if sl.Erase(100) {
		t.Error("Erase should return false for non-existent element 100")
	}

	// Test 6: Erase all elements and verify Skiplist is empty.
	for _, num := range []int{1, 5, 7, 9, 12} {
		if !sl.Erase(num) {
			t.Errorf("Erase should return true for existing element %d", num)
		}
	}
	for _, num := range []int{1, 5, 7, 9, 12} {
		if sl.Search(num) {
			t.Errorf("After erasing, search for %d should return false", num)
		}
	}
}
