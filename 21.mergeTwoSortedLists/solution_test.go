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
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return dummyHead.Next
}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(
		t,
		&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		mergeTwoLists(
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		),
	)
	assert.Equal(
		t,
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}},
		mergeTwoLists(
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{3, &ListNode{4, nil}},
		),
	)
}
