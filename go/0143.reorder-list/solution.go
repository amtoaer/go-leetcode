// Created by amtoaer at 2025/11/20 13:49
// leetgo: 1.4.15
// https://leetcode.cn/problems/reorder-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func splitList(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head.Next
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}
	next := slow.Next
	slow.Next = nil
	return head, next
}

func reverseList(head *ListNode) *ListNode {
	var next, prev *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func mergeList(left, right *ListNode) {
	dummy := &ListNode{}
	cur := dummy
	takeLeft := true
	for left != nil && right != nil {
		if takeLeft {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
		takeLeft = !takeLeft
	}
	if left != nil {
		cur.Next = left
	} else {
		cur.Next = right
	}
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	left, right := splitList(head)
	right = reverseList(right)
	mergeList(left, right)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	reorderList(head)
	ans := head

	fmt.Println("\noutput:", Serialize(ans))
}
