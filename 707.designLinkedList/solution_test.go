package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] Design Linked List
 */

// @lc code=start

type Node struct {
	Val        int
	prev, next *Node
}

type MyLinkedList struct {
	start, end *Node
	size       int
}

func Constructor() MyLinkedList {
	start, end := &Node{}, &Node{}
	start.next, end.prev = end, start
	return MyLinkedList{
		start, end, 0,
	}
}

func (m *MyLinkedList) getNode(index int) *Node {
	if index >= m.size {
		return nil
	}
	if index > m.size>>1 {
		cur := m.end
		for i := 0; i < m.size-index; i++ {
			cur = cur.prev
		}
		return cur
	}
	cur := m.start
	for i := 0; i <= index; i++ {
		cur = cur.next
	}
	return cur
}

func (m *MyLinkedList) Get(index int) int {
	if n := m.getNode(index); n != nil {
		return n.Val
	}
	return -1
}

func (m *MyLinkedList) AddAtHead(val int) {
	prev, next := m.start, m.start.next
	newNode := &Node{Val: val, prev: prev, next: next}
	prev.next, next.prev = newNode, newNode
	m.size++
}

func (m *MyLinkedList) AddAtTail(val int) {
	prev, next := m.end.prev, m.end
	newNode := &Node{Val: val, prev: prev, next: next}
	prev.next, next.prev = newNode, newNode
	m.size++
}

func (m *MyLinkedList) AddAtIndex(index int, val int) {
	if index > m.size {
		return
	}
	newNode := &Node{Val: val}
	var prev, next *Node
	if n := m.getNode(index); n != nil {
		prev, next = n.prev, n
	} else {
		prev, next = m.end.prev, m.end
	}
	newNode.prev, newNode.next = prev, next
	prev.next, next.prev = newNode, newNode
	m.size++
}

func (m *MyLinkedList) DeleteAtIndex(index int) {
	if n := m.getNode(index); n != nil {
		prev, next := n.prev, n.next
		prev.next, next.prev = next, prev
		m.size--
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

func Test(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtHead(2)
	obj.AddAtHead(3)
	obj.AddAtTail(4)
	obj.AddAtTail(5)
	obj.AddAtIndex(1, 10)
	assert.Equal(t, obj.Get(0), 3)
	assert.Equal(t, obj.Get(1), 10)
	assert.Equal(t, obj.Get(2), 2)
	assert.Equal(t, obj.Get(3), 1)
}
