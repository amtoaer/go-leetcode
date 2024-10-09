package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
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
func dfs(root *TreeNode) []int {
	var res []int
	var dfsHelper func(node *TreeNode)
	dfsHelper = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
		}
		dfsHelper(node.Left)
		dfsHelper(node.Right)
	}
	dfsHelper(root)
	return res
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(dfs(root1), dfs(root2))
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input1 *TreeNode
		input2 *TreeNode
		want   bool
	}{{nil, nil, true}}
	for _, tt := range tc {
		assert.Equal(t, tt.want, leafSimilar(tt.input1, tt.input2))
	}
}
