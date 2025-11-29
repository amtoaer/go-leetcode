// Created by amtoaer at 2025/11/29 17:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Next.Val == cur.Val {
				cur.Next = cur.Next.Next
			}
		}
		cur = cur.Next
	}
	return head
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := deleteDuplicates(head)

	fmt.Println("\noutput:", Serialize(ans))
}
