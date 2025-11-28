// Created by amtoaer at 2025/11/28 21:38
// leetgo: 1.4.15
// https://leetcode.cn/problems/diameter-of-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := maxDepth(node.Left), maxDepth(node.Right)
		ans = max(ans, left+right)
		return max(left, right) + 1
	}
	maxDepth(root)
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := diameterOfBinaryTree(root)

	fmt.Println("\noutput:", Serialize(ans))
}
