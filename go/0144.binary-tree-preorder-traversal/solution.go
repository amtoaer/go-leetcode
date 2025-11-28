// Created by amtoaer at 2025/11/28 11:24
// leetgo: 1.4.15
// https://leetcode.cn/problems/binary-tree-preorder-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func preorderTraversal(root *TreeNode) (ans []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := preorderTraversal(root)

	fmt.Println("\noutput:", Serialize(ans))
}
