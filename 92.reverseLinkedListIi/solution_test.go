package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=20004
 *
 * [92] 反转链表 II
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if left == right {
		return head
	}
	var prev *ListNode
	cur := head
	cnt := 1
	for cnt != left {
		prev = cur
		cur = cur.Next
		cnt++
	}
	before := prev
	prev = cur
	cur = cur.Next
	leftNode := prev
	for cnt != right {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		cnt++
	}
	leftNode.Next = cur
	if before != nil {
		before.Next = prev
		return head
	}
	return prev
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/

func arrayToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func listToArray(head *ListNode) []int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		head     []int
		left     int
		right    int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{[]int{5}, 1, 1, []int{5}},
		{[]int{1, 2, 3, 4, 5}, 1, 5, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, 3, 4, []int{1, 2, 4, 3, 5}},
		{[]int{1, 2}, 1, 2, []int{2, 1}},
	}

	for _, test := range tests {
		head := arrayToList(test.head)
		result := reverseBetween(head, test.left, test.right)
		resultArray := listToArray(result)
		for i, v := range resultArray {
			if v != test.expected[i] {
				t.Errorf("For input %v, left %d, right %d, expected %v but got %v", test.head, test.left, test.right, test.expected, resultArray)
				break
			}
		}
	}
}
