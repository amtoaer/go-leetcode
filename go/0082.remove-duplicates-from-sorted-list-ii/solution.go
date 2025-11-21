// Created by amtoaer at 2025/11/21 17:52
// leetgo: 1.4.15
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		start := cur.Next
		if start.Next == nil || start.Next.Val != start.Val {
			cur = cur.Next
		} else {
			for cur.Next != nil && cur.Next.Val == start.Val {
				cur.Next = cur.Next.Next
			}
		}
	}
	return dummy.Next
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := deleteDuplicates(head)

	fmt.Println("\noutput:", Serialize(ans))
}
