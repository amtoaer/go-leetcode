/*
 * @lc app=leetcode.cn id=257 lang=golang
 * @lcpr version=30104
 *
 * [257] 二叉树的所有路径
 */

package main

import (
	"strconv"
	"strings"
	"testing"
)

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
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var binaryTreePathsHelper func(*TreeNode, []string)
	binaryTreePathsHelper = func(node *TreeNode, items []string) {
		if node == nil {
			return
		}
		items = append(items, strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			res = append(res, strings.Join(items, "->"))
			return
		}
		binaryTreePathsHelper(node.Left, items)
		binaryTreePathsHelper(node.Right, items)
	}
	binaryTreePathsHelper(root, []string{})
	return res
}

// @lc code=end

func Test(t *testing.T) {
	// Create first test tree: [1,2,3,null,5]
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// Create second test tree: [1]
	tree2 := &TreeNode{
		Val: 1,
	}

	// Create third test tree: a more complex binary tree
	tree3 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	tests := []struct {
		name     string
		root     *TreeNode
		expected []string
	}{
		{
			name:     "Tree with left and right children",
			root:     tree1,
			expected: []string{"1->2->5", "1->3"},
		},
		{
			name:     "Tree with only root",
			root:     tree2,
			expected: []string{"1"},
		},
		{
			name:     "Complete binary tree",
			root:     tree3,
			expected: []string{"4->2->1", "4->2->3", "4->7->6", "4->7->9"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binaryTreePaths(tt.root)
			if len(got) != len(tt.expected) {
				t.Errorf("binaryTreePaths() length = %v, expected length %v", len(got), len(tt.expected))
				return
			}

			// Create a map to check if all expected paths are present
			expectedMap := make(map[string]bool)
			for _, path := range tt.expected {
				expectedMap[path] = true
			}

			for _, path := range got {
				if !expectedMap[path] {
					t.Errorf("binaryTreePaths() unexpected path = %v", path)
				}
			}
		})
	}
}

/*
// @lcpr case=start
// [1,2,3,null,5]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
