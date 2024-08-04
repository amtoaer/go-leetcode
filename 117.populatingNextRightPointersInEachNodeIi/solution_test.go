package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, curNode *Node
		innerConnect := func(n *Node) {
			if n == nil {
				return
			}
			if nextStart == nil {
				nextStart = n
				curNode = n
				return
			}
			curNode.Next = n
			curNode = n
		}
		for n := start; n != nil; n = n.Next {
			innerConnect(n.Left)
			innerConnect(n.Right)
		}
		start = nextStart
	}
	return root
}

// @lc code=end

func Test(t *testing.T) {
	assert.Nil(t, connect(nil))
}
