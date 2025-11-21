// Created by amtoaer at 2025/11/21 15:30
// leetgo: 1.4.15
// https://leetcode.cn/problems/reverse-nodes-in-k-group/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for {
		cur = reverseNextK(cur, k)
		if cur == nil {
			return dummy.Next
		}
	}
}

func reverseNextK(node *ListNode, k int) *ListNode {
	last := node
	// 确保有足够的 node
	for range k {
		if last.Next == nil {
			return nil
		}
		last = last.Next
	}
	start, end := node.Next, last.Next
	prev, cur := end, start
	for cur != end {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	node.Next = prev
	return start
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := reverseKGroup(head, k)

	fmt.Println("\noutput:", Serialize(ans))
}
