package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=21 lang=golang
 * @lcpr version=20004
 *
 * [21] 合并两个有序链表
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	left, right := list1, list2
	for left != nil && right != nil {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}
	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			list1: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			list2: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			want:  &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			list1: nil,
			list2: nil,
			want:  nil,
		},
		{
			list1: nil,
			list2: &ListNode{0, nil},
			want:  &ListNode{0, nil},
		},
		{
			list1: &ListNode{2, nil},
			list2: &ListNode{1, nil},
			want:  &ListNode{1, &ListNode{2, nil}},
		},
	}

	for _, tt := range tests {
		got := mergeTwoLists(tt.list1, tt.list2)
		for got != nil && tt.want != nil {
			if got.Val != tt.want.Val {
				t.Errorf("mergeTwoLists() = %v, want %v", got.Val, tt.want.Val)
			}
			got = got.Next
			tt.want = tt.want.Next
		}
		if got != nil || tt.want != nil {
			t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
		}
	}
}
