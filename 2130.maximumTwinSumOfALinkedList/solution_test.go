package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=2130 lang=golang
 *
 * [2130] Maximum Twin Sum of a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	var nums []int
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast.Next != nil {
		slow = slow.Next
		nums = append(nums, slow.Val)
		// 题目确保长度是偶数，直接这样写是安全的
		fast = fast.Next.Next
	}
	res := math.MinInt32
	idx := len(nums) - 1
	slow = slow.Next
	for slow != nil {
		res = max(nums[idx]+slow.Val, res)
		idx--
		slow = slow.Next
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{[]int{5, 4, 2, 1}, 6},
		{[]int{4, 2, 2, 3}, 7},
	}
	for _, tt := range tc {
		var head, cur *ListNode
		for i, v := range tt.input {
			if i == 0 {
				head = &ListNode{Val: v}
				cur = head
			} else {
				cur.Next = &ListNode{Val: v}
				cur = cur.Next
			}
		}
		assert.Equal(t, tt.output, pairSum(head))
	}
}
