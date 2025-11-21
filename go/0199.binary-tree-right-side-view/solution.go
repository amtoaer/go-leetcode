// Created by amtoaer at 2025/11/21 14:12
// leetgo: 1.4.15
// https://leetcode.cn/problems/binary-tree-right-side-view/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		ans = append(ans, queue[len(queue)-1].Val)
		queueLen := len(queue)
		for _, node := range queue[:queueLen] {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[queueLen:]
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := rightSideView(root)

	fmt.Println("\noutput:", Serialize(ans))
}
