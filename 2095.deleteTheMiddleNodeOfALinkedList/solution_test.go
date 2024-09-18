package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=2095 lang=golang
 *
 * [2095] Delete the Middle Node of a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var prev *ListNode
	slow, fast := head, head
	for fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	if prev.Next != nil {
		prev.Next = slow.Next
	}
	return head
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *ListNode
		output *ListNode
	}{
		{nil, nil},
		{&ListNode{Val: 1}, nil},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, deleteMiddle(tt.input))
	}
}
