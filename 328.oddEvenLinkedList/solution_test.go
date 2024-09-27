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
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	tmp1, tmp2 := dummy1, dummy2
	var flag bool
	for cur := head; cur != nil; cur = cur.Next {
		if !flag {
			tmp1.Next = cur
			tmp1 = tmp1.Next
		} else {
			tmp2.Next = cur
			tmp2 = tmp2.Next
		}
		flag = !flag
	}
	tmp1.Next = dummy2.Next
	tmp2.Next = nil
	return dummy1.Next
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  *ListNode
		output *ListNode
	}{
		{nil, nil},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, oddEvenList(tt.input))
	}
}
