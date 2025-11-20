// Created by amtoaer at 2025/11/18 17:02
// leetgo: 1.4.15
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	}
	return right
}

// @lc code=end
