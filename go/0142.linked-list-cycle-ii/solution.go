// Created by amtoaer at 2025/11/22 09:45
// leetgo: 1.4.15
// https://leetcode.cn/problems/linked-list-cycle-ii/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := make(map[*ListNode]any)
	for cur := head; cur.Next != nil; cur = cur.Next {
		if _, ok := m[cur]; ok {
			return cur
		}
		m[cur] = struct{}{}
	}
	return nil
}

// @lc code=end
