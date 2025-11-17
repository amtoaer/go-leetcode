// Created by amtoaer at 2025/11/18 02:38
// leetgo: 1.4.15
// https://leetcode.cn/problems/binary-tree-level-order-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func levelOrder(root *TreeNode) (ans [][]int) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		beforeLen := len(queue)
		var levelAns []int
		for _, node := range queue[:beforeLen] {
			if node != nil {
				levelAns = append(levelAns, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[beforeLen:]
		if len(levelAns) > 0 {
			ans = append(ans, levelAns)
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := levelOrder(root)

	fmt.Println("\noutput:", Serialize(ans))
}
