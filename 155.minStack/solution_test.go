package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	if len(m.minStack) == 0 {
		m.minStack = append(m.minStack, val)
	} else {
		m.minStack = append(m.minStack, min(val, m.minStack[len(m.minStack)-1]))
	}
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
	m.minStack = m.minStack[:len(m.minStack)-1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

func Test(t *testing.T) {
	stack := Constructor()
	stack.Push(3)
	stack.Push(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(4)
	stack.Push(6)
	assert.Equal(t, 6, stack.Top())
	assert.Equal(t, 1, stack.GetMin())
}
