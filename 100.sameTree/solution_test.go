package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=100 lang=golang
 * @lcpr version=20004
 *
 * [100] 相同的树
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n[1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[1,null,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1]\n[1,1,2]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want: true,
		},
		{
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			q:    &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			want: false,
		},
		{
			p:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			q:    &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			want: false,
		},
		{
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			p:    &TreeNode{Val: 1},
			q:    &TreeNode{Val: 1},
			want: true,
		},
		{
			p:    &TreeNode{Val: 1},
			q:    &TreeNode{Val: 2},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isSameTree(tt.p, tt.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
