// Created by amtoaer at 2025/11/20 11:36
// leetgo: 1.4.15
// https://leetcode.cn/problems/reverse-linked-list-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	var before, prev, next *ListNode
	step := right - left + 1
	for before = dummy; left > 1; left-- {
		before = before.Next
	}
	start := before.Next
	cur := start
	for step > 0 {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		step--
	}
	before.Next = prev
	start.Next = cur
	return dummy.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	left := Deserialize[int](ReadLine(stdin))
	right := Deserialize[int](ReadLine(stdin))
	ans := reverseBetween(head, left, right)

	fmt.Println("\noutput:", Serialize(ans))
}
