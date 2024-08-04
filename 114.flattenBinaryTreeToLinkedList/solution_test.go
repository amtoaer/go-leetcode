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
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
func flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
			continue
		}
		next := cur.Left
		pre := next
		for pre.Right != nil {
			pre = pre.Right
		}
		pre.Right = cur.Right
		cur.Left, cur.Right = nil, next
	}
}

// @lc code=end

func Test(t *testing.T) {
	// 测试用例
	testCases := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{
								Val: 5,
								Right: &TreeNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		flatten(tc.root)
		assert.True(t, reflect.DeepEqual(tc.want, tc.root))
	}
}
