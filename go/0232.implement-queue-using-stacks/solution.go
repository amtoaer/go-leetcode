// Created by amtoaer at 2025/11/22 11:07
// leetgo: 1.4.15
// https://leetcode.cn/problems/implement-queue-using-stacks/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MyQueue struct {
	stack1, stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (m *MyQueue) Push(x int) {
	for len(m.stack2) != 0 {
		m.stack1 = append(m.stack1, m.stack2[len(m.stack2)-1])
		m.stack2 = m.stack2[:len(m.stack2)-1]
	}
	m.stack1 = append(m.stack1, x)
}

func (m *MyQueue) Pop() (ans int) {
	for len(m.stack1) != 0 {
		m.stack2 = append(m.stack2, m.stack1[len(m.stack1)-1])
		m.stack1 = m.stack1[:len(m.stack1)-1]
	}
	ans = m.stack2[len(m.stack2)-1]
	m.stack2 = m.stack2[:len(m.stack2)-1]
	return
}

func (m *MyQueue) Peek() (ans int) {
	for len(m.stack1) != 0 {
		m.stack2 = append(m.stack2, m.stack1[len(m.stack1)-1])
		m.stack1 = m.stack1[:len(m.stack1)-1]
	}
	ans = m.stack2[len(m.stack2)-1]
	return
}

func (m *MyQueue) Empty() bool {
	return len(m.stack1) == 0 && len(m.stack2) == 0
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
			x := Deserialize[int](methodParams[0])
			obj.Push(x)
			output = append(output, "null")
		case "pop":
			ans := Serialize(obj.Pop())
			output = append(output, ans)
		case "peek":
			ans := Serialize(obj.Peek())
			output = append(output, ans)
		case "empty":
			ans := Serialize(obj.Empty())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
