// Created by amtoaer at 2025/11/21 13:12
// leetgo: 1.4.15
// https://leetcode.cn/problems/intersection-of-two-linked-lists/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}
		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}

// @lc code=end
