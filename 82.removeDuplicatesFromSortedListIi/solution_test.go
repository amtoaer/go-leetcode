package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=82 lang=golang
 * @lcpr version=20004
 *
 * [82] 删除排序链表中的重复元素 II
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Next != nil && cur.Next.Val == cur.Next.Next.Val {
			val := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur = cur.Next
			}
			continue
		}
		cur = cur.Next
		prev.Next = cur
		prev = prev.Next
	}
	prev.Next = nil
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,3,4,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{
			input:    createList([]int{1, 2, 3, 3, 4, 4, 5}),
			expected: createList([]int{1, 2, 5}),
		},
		{
			input:    createList([]int{1, 1, 1, 2, 3}),
			expected: createList([]int{2, 3}),
		},
		{
			input:    createList([]int{1, 1, 2, 2, 3, 3}),
			expected: createList([]int{}),
		},
		{
			input:    createList([]int{1, 2, 3, 4, 5}),
			expected: createList([]int{1, 2, 3, 4, 5}),
		},
		{
			input:    createList([]int{1, 1}),
			expected: createList([]int{}),
		},
	}

	for _, test := range tests {
		result := deleteDuplicates(test.input)
		if !compareLists(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", listToSlice(test.input), listToSlice(test.expected), listToSlice(result))
		}
	}
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
