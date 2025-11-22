// Created by amtoaer at 2025/11/22 11:32
// leetgo: 1.4.15
// https://leetcode.cn/problems/binary-tree-inorder-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func inorderTraversal(root *TreeNode) (ans []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		ans = append(ans, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := inorderTraversal(root)

	fmt.Println("\noutput:", Serialize(ans))
}
