package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type Node struct {
	Key, Val int
	pre      *Node
	next     *Node
}

type LRUCache struct {
	capacity   int
	head, tail *Node
	m          map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		m:        make(map[int]*Node),
	}
}

func (l *LRUCache) moveToFirst(n *Node) {
	pre, next := n.pre, n.next
	if pre != nil && next != nil {
		pre.next, next.pre = next, pre
	}
	pre, next = l.head, l.head.next
	pre.next, next.pre = n, n
	n.pre, n.next = pre, next
}

func (l *LRUCache) removeLast() *Node {
	pre, cur, next := l.tail.pre.pre, l.tail.pre, l.tail
	pre.next, next.pre = next, pre
	return cur
}

func (l *LRUCache) Get(key int) int {
	if n, ok := l.m[key]; ok {
		l.moveToFirst(n)
		return n.Val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if n, ok := l.m[key]; ok {
		l.moveToFirst(n)
		n.Val = value
		return
	}
	newNode := &Node{
		Key: key,
		Val: value,
	}
	l.moveToFirst(newNode)
	l.m[key] = newNode
	if len(l.m) > l.capacity {
		n := l.removeLast()
		delete(l.m, n.Key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func Test(t *testing.T) {
	// Input
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	// Output
	// [null, null, null, 1, null, -1, null, -1, 3, 4]

	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	if obj.Get(1) != 1 {
		t.Errorf("Get(1) failed")
	}
	obj.Put(3, 3)
	if obj.Get(2) != -1 {
		t.Errorf("Get(2) failed")
	}
	obj.Put(4, 4)
	if obj.Get(1) != -1 {
		t.Errorf("Get(1) failed")
	}
	if obj.Get(3) != 3 {
		t.Errorf("Get(3) failed")
	}
	if obj.Get(4) != 4 {
		t.Errorf("Get(4) failed")
	}
}
