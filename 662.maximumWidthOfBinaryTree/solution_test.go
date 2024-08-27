package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] Maximum Width of Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	queue := []Pair{{root, 1}}
	for len(queue) > 0 {
		lenQueue := len(queue)
		res = max(res, queue[lenQueue-1].index-queue[0].index+1)
		for _, p := range queue[:lenQueue] {
			if p.node.Left != nil {
				queue = append(queue, Pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				queue = append(queue, Pair{p.node.Right, p.index*2 + 1})
			}
		}
		queue = queue[lenQueue:]
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		root   *TreeNode
		expect int
	}{
		{nil, 0},
		{&TreeNode{1, nil, nil}, 1},
	}
	for _, tt := range tc {
		assert.Equal(t, widthOfBinaryTree(tt.root), tt.expect)
	}
}
