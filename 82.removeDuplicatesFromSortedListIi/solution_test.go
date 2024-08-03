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
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{0, head}
	head = dummy
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			reVal := head.Next.Val
			for head.Next != nil && head.Next.Val == reVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}

// @lc code=end

func Test(t *testing.T) {
	// test cases
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{&ListNode{1, &ListNode{1, &ListNode{1, nil}}}, nil},
		{&ListNode{1, &ListNode{1, &ListNode{2, nil}}}, &ListNode{2, nil}},
		{&ListNode{1, &ListNode{2, &ListNode{2, nil}}}, &ListNode{1, nil}},
	}

	compareList := func(l1, l2 *ListNode) bool {
		for l1 != nil && l2 != nil {
			if l1.Val != l2.Val {
				return false
			}
			l1 = l1.Next
			l2 = l2.Next
		}
		return l1 == nil && l2 == nil
	}

	for _, c := range cases {
		assert.True(t, compareList(deleteDuplicates(c.input), c.expect))
	}
}
