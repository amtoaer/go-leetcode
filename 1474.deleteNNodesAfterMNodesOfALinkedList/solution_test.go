/*
 * @lc app=leetcode.cn id=1474 lang=golang
 * @lcpr version=30202
 *
 * [1474] 删除链表 M 个节点之后的 N 个节点
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
func deleteNodes(head *ListNode, m int, n int) *ListNode {
	cur := head
	for cur != nil {
		for i := 1; cur.Next != nil && i < m; i++ {
			cur = cur.Next
		}
		tmp := cur
		for i := 0; tmp != nil && i <= n; i++ {
			tmp = tmp.Next
			cur.Next = tmp
		}
		cur = cur.Next
	}
	return head
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		head *ListNode
		m    int
		n    int
		want *ListNode
	}{
		{
			&ListNode{9, &ListNode{3, &ListNode{7, &ListNode{7, &ListNode{9, &ListNode{10, &ListNode{8, &ListNode{2, nil}}}}}}}},
			1, 2,
			&ListNode{9, &ListNode{7, &ListNode{8, nil}}},
		},
	}
	for _, tt := range tc {
		result := deleteNodes(tt.head, tt.m, tt.n)
		// Compare linked lists by converting to slices
		var resultSlice, wantSlice []int
		for cur := result; cur != nil; cur = cur.Next {
			resultSlice = append(resultSlice, cur.Val)
		}
		for cur := tt.want; cur != nil; cur = cur.Next {
			wantSlice = append(wantSlice, cur.Val)
		}
		assert.Equal(t, wantSlice, resultSlice)
	}
}

/*
// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10,11,12,13]\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10,11]\n1\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10,11]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// [9,3,7,7,9,10,8,2]\n1\n2\n
// @lcpr case=end

*/
