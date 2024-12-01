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
 * @lcpr version=20004
 *
 * [61] 旋转链表
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	cnt := 1
	for cur.Next != nil {
		cur = cur.Next
		cnt++
	}
	end := cur
	k %= cnt
	if k == 0 {
		return head
	}
	cur = head
	for i := 0; i < cnt-k-1; i++ {
		cur = cur.Next
	}
	tmp := cur.Next
	cur.Next = nil
	end.Next = head
	return tmp
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2]\n4\n
// @lcpr case=end

*/

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for _, num := range nums[1:] {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return head
}

func Test(t *testing.T) {
	tests := []struct {
		input  []int
		k      int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1, 2}, 1, []int{2, 1}},
		{[]int{1}, 99, []int{1}},
	}

	for _, test := range tests {
		head := sliceToList(test.input)
		expected := sliceToList(test.output)
		result := rotateRight(head, test.k)
		if !equal(result, expected) {
			t.Errorf("rotateRight(%v, %d) = %v; want %v", test.input, test.k, listToSlice(result), test.output)
		}
	}
}

func equal(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
