/*
 * @lc app=leetcode.cn id=1265 lang=golang
 * @lcpr version=30202
 *
 * [1265] 逆序打印不可变链表
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var prints []int

type ListNode struct {
	val  int
	next *ListNode
}

func (l *ListNode) getNext() ImmutableListNode {
	return l.next
}

func (l *ListNode) printValue() {
	prints = append(prints, l.val)
}

type ImmutableListNode = *ListNode

// @lc code=start
/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */

func printLinkedListInReverse(head ImmutableListNode) {
	for current := head; current != nil; current = current.getNext() {
		defer current.printValue()
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{0, -4, -1, 3, -5}, []int{-5, 3, -1, -4, 0}},
		{[]int{-2, 0, 6, 4, 4, -6}, []int{-6, 4, 4, 6, 0, -2}},
	}
	for _, tt := range tc {
		prints = []int{}
		var head *ListNode
		if len(tt.input) > 0 {
			head = &ListNode{val: tt.input[0]}
			current := head
			for i := 1; i < len(tt.input); i++ {
				current.next = &ListNode{val: tt.input[i]}
				current = current.next
			}
		}
		printLinkedListInReverse(head)
		assert.Equal(t, tt.want, prints)
	}
}

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [0,-4,-1,3,-5]\n
// @lcpr case=end

// @lcpr case=start
// [-2,0,6,4,4,-6]\n
// @lcpr case=end

*/
