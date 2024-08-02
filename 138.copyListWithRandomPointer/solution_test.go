package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	// 将复制的节点接到当前节点的下一个节点
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{
			Val:    cur.Val,
			Next:   next,
			Random: cur.Random,
		}
		cur = next
	}
	cur = head
	for cur != nil {
		if cur.Next.Random != nil {
			cur.Next.Random = cur.Next.Random.Next
		}
		cur = cur.Next.Next
	}
	res := head.Next
	cur = head
	for cur != nil {
		tmp := cur.Next
		next := tmp.Next
		cur.Next = next
		if next != nil {
			tmp.Next = next.Next
		}
		cur = next
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	assert.Nil(t, copyRandomList(nil))
}
