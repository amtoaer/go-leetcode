// Created by amtoaer at 2025/11/22 09:45
// leetgo: 1.4.15
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for range n {
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return dummy.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := removeNthFromEnd(head, n)

	fmt.Println("\noutput:", Serialize(ans))
}
