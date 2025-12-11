// Created by amtoaer at 2025/12/12 00:59
// leetgo: 1.4.15
// https://leetcode.cn/problems/copy-list-with-random-pointer/

package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=begin

func copyRandomList(head *Node) (ans *Node) {
	if head == nil {
		return nil
	}
	var tail *Node
	var length int
	for cur := head; cur != nil; cur = cur.Next {
		tail = cur
		length++
	}
	old, new := head, tail
	for old != tail.Next {
		new.Next = &Node{Val: old.Val}
		old = old.Next
		new = new.Next
	}
	old, new = head, tail.Next
	for new != nil {
		if old.Random != nil {
			new.Random = old.Random
			for range length {
				new.Random = new.Random.Next
			}
		}
		new = new.Next
		old = old.Next
	}
	ans = tail.Next
	tail.Next = nil
	return
}

// @lc code=end
