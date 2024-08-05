package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		root *TreeNode
		want int
	}{
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 3},
		{
			&TreeNode{
				1,
				&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
				&TreeNode{3, &TreeNode{6, nil, nil}, nil},
			},
			6,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, countNodes(tt.root))
	}
}
