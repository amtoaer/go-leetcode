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
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for cur := headA; cur != nil; cur = cur.Next {
		m[cur] = struct{}{}
	}
	for cur := headB; cur != nil; cur = cur.Next {
		if _, ok := m[cur]; ok {
			return cur
		}
	}
	return nil
}

// @lc code=end

func Test(t *testing.T) {
	assert.Nil(t, getIntersectionNode(nil, nil))
}
