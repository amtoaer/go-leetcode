// Created by amtoaer at 2025/11/27 18:23
// leetgo: 1.4.15
// https://leetcode.cn/problems/min-stack/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MinStack struct {
	stack, minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	if len(m.minStack) > 0 {
		m.minStack = append(m.minStack, min(m.minStack[len(m.minStack)-1], val))
	} else {
		m.minStack = append(m.minStack, val)
	}
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
	m.minStack = m.minStack[:len(m.minStack)-1]
}

func (m *MinStack) Top() (ans int) {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "push":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			obj.Push(val)
			output = append(output, "null")
		case "pop":
			obj.Pop()
			output = append(output, "null")
		case "top":
			ans := Serialize(obj.Top())
			output = append(output, ans)
		case "getMin":
			ans := Serialize(obj.GetMin())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
