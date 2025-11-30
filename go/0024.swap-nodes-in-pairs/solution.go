// Created by amtoaer at 2025/11/30 15:36
// leetgo: 1.4.15
// https://leetcode.cn/problems/swap-nodes-in-pairs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		next := cur.Next.Next.Next
		tmp := cur.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmp
		tmp.Next = next
		cur = tmp
	}
	return dummy.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := swapPairs(head)

	fmt.Println("\noutput:", Serialize(ans))
}
