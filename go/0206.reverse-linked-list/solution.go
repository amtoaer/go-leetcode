// Created by amtoaer at 2025/11/17 23:23
// leetgo: 1.4.15
// https://leetcode.cn/problems/reverse-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := reverseList(head)

	fmt.Println("\noutput:", Serialize(ans))
}
