// Created by amtoaer at 2025/12/18 22:32
// leetgo: 1.4.15
// https://leetcode.cn/problems/binary-tree-postorder-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func postorderTraversal(root *TreeNode) (ans []int) {
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			helper(node.Left)
		}
		if node.Right != nil {
			helper(node.Right)
		}
		ans = append(ans, node.Val)
	}
	helper(root)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := postorderTraversal(root)

	fmt.Println("\noutput:", Serialize(ans))
}
