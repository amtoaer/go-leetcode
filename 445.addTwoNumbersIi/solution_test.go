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
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] Add Two Numbers II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverseList(l1), reverseList(l2)
	var sum int
	dummyHead := &ListNode{}
	cur := dummyHead
	for l1 != nil || l2 != nil || sum != 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		sum /= 10
		cur = cur.Next
	}
	if dummyHead.Next == nil {
		return nil
	}
	return reverseList(dummyHead.Next)
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(
		t,
		&ListNode{
			Val:  7,
			Next: &ListNode{Val: 8, Next: &ListNode{Val: 0, Next: &ListNode{Val: 7}}},
		},
		addTwoNumbers(
			&ListNode{
				Val:  7,
				Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
		),
	)
	assert.Equal(t, &ListNode{Val: 0}, addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 0}))
}
