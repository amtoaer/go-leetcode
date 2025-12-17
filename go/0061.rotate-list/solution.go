// Created by amtoaer at 2025/12/12 15:15
// leetgo: 1.4.15
// https://leetcode.cn/problems/rotate-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func rotateRight(head *ListNode, k int) *ListNode {
	var length int
	var tail *ListNode
	for cur := head; cur != nil; cur = cur.Next {
		tail = cur
		length++
	}
	if length == 0 {
		return head
	}
	k %= length
	if k == 0 {
		return head
	}
	prev := head
	for range length - k - 1 {
		prev = prev.Next
	}
	ans := prev.Next
	prev.Next = nil
	tail.Next = head
	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := rotateRight(head, k)

	fmt.Println("\noutput:", Serialize(ans))
}
