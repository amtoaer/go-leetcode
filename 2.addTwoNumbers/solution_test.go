package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=20004
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var left int
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil || left != 0 {
		if l1 != nil {
			left += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			left += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{
			Val: left % 10,
		}
		cur = cur.Next
		left /= 10
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/

func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(l *ListNode) []int {
	var result []int
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}

func Test(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{[]int{1, 8}, []int{0}, []int{1, 8}},
		{[]int{5}, []int{5}, []int{0, 1}},
	}

	for _, test := range tests {
		l1 := createList(test.l1)
		l2 := createList(test.l2)
		result := addTwoNumbers(l1, l2)
		resultSlice := listToSlice(result)
		if !equal(resultSlice, test.expected) {
			t.Errorf("addTwoNumbers(%v, %v) = %v; expected %v", test.l1, test.l2, resultSlice, test.expected)
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
