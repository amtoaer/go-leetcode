// Created by amtoaer at 2025/11/27 15:58
// leetgo: 1.4.15
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	i := 0
	for i < len(inorder) {
		if inorder[i] == preorder[0] {
			break
		}
		i++
	}
	node.Left = buildTree(preorder[1:i+1], inorder[:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return node
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	preorder := Deserialize[[]int](ReadLine(stdin))
	inorder := Deserialize[[]int](ReadLine(stdin))
	ans := buildTree(preorder, inorder)

	fmt.Println("\noutput:", Serialize(ans))
}
