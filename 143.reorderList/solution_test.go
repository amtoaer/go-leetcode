package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=143 lang=golang
 * @lcpr version=20004
 *
 * [143] 重排链表
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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	mid := slow
	var prev *ListNode
	cur := mid.Next
	mid.Next = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	left := head
	right := prev
	for right != nil {
		tmp := left.Next
		left.Next = right
		left = tmp
		tmp = right.Next
		right.Next = left
		right = tmp
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4}, expected: []int{1, 4, 2, 3}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 5, 2, 4, 3}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{1, 2}, expected: []int{1, 2}},
		{input: []int{1, 2, 3}, expected: []int{1, 3, 2}},
	}

	for _, test := range tests {
		head := createList(test.input)
		reorderList(head)
		result := listToSlice(head)
		if !equal(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
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

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
