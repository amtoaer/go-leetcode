// Created by amtoaer at 2025/11/20 13:41
// leetgo: 1.4.15
// https://leetcode.cn/problems/linked-list-cycle/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// @lc code=end
