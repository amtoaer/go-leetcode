/*
 * @lc app=leetcode.cn id=1290 lang=golang
 * @lcpr version=30201
 *
 * [1290] 二进制链表转整数
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var res int
	for head != nil {
		res = res*2 + head.Val
		head = head.Next
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		head *ListNode
		want int
	}{
		{&ListNode{1, &ListNode{0, &ListNode{1, nil}}}, 5},
		{&ListNode{0, nil}, 0},
		{&ListNode{1, nil}, 1},
		{&ListNode{1, &ListNode{0, &ListNode{0, &ListNode{1, &ListNode{0, &ListNode{0, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, nil}}}}}}}}}}}}}}}, 18880},
		{&ListNode{0, &ListNode{0, nil}}, 0},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, getDecimalValue(tt.head))
	}
}

/*
// @lcpr case=start
// [1,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,0]\n
// @lcpr case=end

*/
