package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cur, cnt := head, 0
	for cur != nil {
		cur = cur.Next
		cnt++
	}
	k %= cnt
	if k == 0 {
		return head
	}
	cur, cnt = head, cnt-k-1
	for cnt > 0 {
		cur = cur.Next
		cnt--
	}
	tmp := cur.Next
	cur.Next = nil
	cur = tmp
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return tmp
}

// @lc code=end

func Test(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	k := 2
	res := rotateRight(head, k)
	vals := []int{4, 5, 1, 2, 3}
	for _, v := range vals {
		if res.Val != v {
			t.Errorf("Expected %d but got %d", v, res.Val)
		}
		res = res.Next
	}
}
