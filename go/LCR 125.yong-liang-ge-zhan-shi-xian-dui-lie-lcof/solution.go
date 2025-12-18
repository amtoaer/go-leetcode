// Created by amtoaer at 2025/12/18 01:21
// leetgo: 1.4.15
// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type CQueue struct {
	queue []int
}

func Constructor() CQueue {
	return CQueue{
		queue: []int{},
	}
}

func (c *CQueue) AppendTail(value int) {
	c.queue = append(c.queue, value)
}

func (c *CQueue) DeleteHead() (ans int) {
	if len(c.queue) == 0 {
		return -1
	}
	h := c.queue[0]
	c.queue = c.queue[1:]
	return h
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
		case "appendTail":
			methodParams := MustSplitArray(params[i])
			value := Deserialize[int](methodParams[0])
			obj.AppendTail(value)
			output = append(output, "null")
		case "deleteHead":
			ans := Serialize(obj.DeleteHead())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
