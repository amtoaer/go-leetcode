/*
 * @lc app=leetcode.cn id=2415 lang=golang
 * @lcpr version=30111
 *
 * [2415] 反转二叉树的奇数层
 */

// @lcpr-template-start
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	x := &TreeNode{Val: 1}
	assert.Equal(t, x, reverseOddLevels(x))
}

// @lcpr-template-end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	even := true
	for len(nodes) > 0 {
		nodesLen := len(nodes)
		for i := 0; i < nodesLen; i++ {
			node := nodes[i]
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		nodes = nodes[nodesLen:]
		if even {
			for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
				nodes[i].Val, nodes[j].Val = nodes[j].Val, nodes[i].Val
			}
		}
		even = !even
	}
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,5,8,13,21,34]\n
// @lcpr case=end

// @lcpr case=start
// [7,13,11]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]\n
// @lcpr case=end

*/
