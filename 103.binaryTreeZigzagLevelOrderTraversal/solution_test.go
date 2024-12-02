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
 * @lc app=leetcode.cn id=103 lang=golang
 * @lcpr version=20004
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var flag bool
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var row []int
		for _, node := range queue[:length] {
			row = append(row, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if flag {
			left, right := 0, len(row)-1
			for left < right {
				row[left], row[right] = row[right], row[left]
				left++
				right--
			}
		}
		res = append(res, row)
		flag = !flag
		queue = queue[length:]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected [][]int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: [][]int{
				{1},
			},
		},
		{
			root:     nil,
			expected: [][]int{},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: [][]int{
				{1},
				{3, 2},
				{4, 5, 6},
			},
		},
	}

	for _, test := range tests {
		result := zigzagLevelOrder(test.root)
		if !equal(result, test.expected) {
			t.Errorf("For root %v, expected %v, but got %v", test.root, test.expected, result)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
