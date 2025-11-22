// Created by amtoaer at 2025/11/22 11:41
// leetgo: 1.4.15
// https://leetcode.cn/problems/sort-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := partition(head)
	left, right = sortList(left), sortList(right)
	return merge(left, right)
}

func partition(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}
	other := slow.Next
	slow.Next = nil
	return head, other
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	} else {
		cur.Next = right
	}
	return dummy.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := sortList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
