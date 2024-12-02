package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=24 lang=golang
 * @lcpr version=20004
 *
 * [24] 两两交换链表中的节点
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var res, prev *ListNode
	cur := head
	for cur != nil && cur.Next != nil {
		if res == nil {
			res = cur.Next
		}
		tmp := cur.Next
		next := tmp.Next
		tmp.Next = cur
		cur.Next = next
		if prev != nil {
			prev.Next = tmp
		}
		prev = cur
		cur = next
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

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

func Test(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4}, expected: []int{2, 1, 4, 3}},
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{1, 2, 3}, expected: []int{2, 1, 3}},
		{input: []int{1, 2}, expected: []int{2, 1}},
	}

	for _, test := range tests {
		head := createList(test.input)
		result := swapPairs(head)
		resultSlice := listToSlice(result)
		if !equal(resultSlice, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, resultSlice)
		}
	}
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
