package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	first, second []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (m *MyStack) Push(x int) {
	m.second = append(m.second, x)
	m.second = append(m.second, m.first...)
	m.first, m.second = m.second, m.first[:0]
}

func (m *MyStack) Pop() int {
	x := m.first[0]
	m.first = m.first[1:]
	return x
}

func (m *MyStack) Top() int {
	return m.first[0]
}

func (m *MyStack) Empty() bool {
	return len(m.first) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

func Test(t *testing.T) {
}
