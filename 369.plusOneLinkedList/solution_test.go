/*
 * @lc app=leetcode.cn id=369 lang=golang
 * @lcpr version=30202
 *
 * [369] 给单链表加一
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
	head = reverse(head)
	carry := 1
	var prev, current *ListNode
	for current = head; current != nil && carry != 0; current = current.Next {
		tmp := current.Val + carry
		current.Val = tmp % 10
		carry = tmp / 10
		prev = current
	}
	if carry != 0 {
		prev.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return reverse(head)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		head *ListNode
		want *ListNode
	}{
		{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, &ListNode{1, &ListNode{2, &ListNode{4, nil}}}},
		{&ListNode{0, nil}, &ListNode{1, nil}},
	}
	for _, tt := range tc {
		result := plusOne(tt.head)
		for result != nil && tt.want != nil {
			assert.Equal(t, tt.want.Val, result.Val)
			result = result.Next
			tt.want = tt.want.Next
		}
		assert.Nil(t, result)
		assert.Nil(t, tt.want)
	}
}

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
