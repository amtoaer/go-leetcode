// Created by amtoaer at 2025/12/18 01:29
// leetgo: 1.4.15
// https://leetcode.cn/problems/subtree-of-another-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if isSameTree(root, subRoot) {
		return true
	}
	if root == nil {
		return false
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	subRoot := Deserialize[*TreeNode](ReadLine(stdin))
	ans := isSubtree(root, subRoot)

	fmt.Println("\noutput:", Serialize(ans))
}
