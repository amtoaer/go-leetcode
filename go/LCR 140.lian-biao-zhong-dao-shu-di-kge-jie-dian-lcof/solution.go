// Created by amtoaer at 2025/11/27 17:22
// leetgo: 1.4.15
// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func trainingPlan(head *ListNode, cnt int) *ListNode {
	slow, fast := head, head
	for range cnt {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	cnt := Deserialize[int](ReadLine(stdin))
	ans := trainingPlan(head, cnt)

	fmt.Println("\noutput:", Serialize(ans))
}
