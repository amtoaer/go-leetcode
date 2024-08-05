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
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
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
type BSTIterator struct {
	items []int
	idx   int
}

func Constructor(root *TreeNode) BSTIterator {
	var items []int
	var init func(*TreeNode)
	init = func(node *TreeNode) {
		if node == nil {
			return
		}
		init(node.Left)
		items = append(items, node.Val)
		init(node.Right)
	}
	init(root)
	return BSTIterator{
		items: items,
	}
}

func (b *BSTIterator) Next() int {
	item := b.items[b.idx]
	b.idx++
	return item
}

func (b *BSTIterator) HasNext() bool {
	return b.idx < len(b.items)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

func Test(t *testing.T) {
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	obj := Constructor(root)
	if obj.Next() != 3 {
		t.Fatal(obj.Next())
	}
	if obj.Next() != 7 {
		t.Fatal(obj.Next())
	}
	if obj.HasNext() != true {
		t.Fatal(obj.HasNext())
	}
	if obj.Next() != 9 {
		t.Fatal(obj.Next())
	}
	if obj.HasNext() != true {
		t.Fatal(obj.HasNext())
	}
	if obj.Next() != 15 {
		t.Fatal(obj.Next())
	}
	if obj.HasNext() != true {
		t.Fatal(obj.HasNext())
	}
	if obj.Next() != 20 {
		t.Fatal(obj.Next())
	}
	if obj.HasNext() != false {
		t.Fatal(obj.HasNext())
	}
}
