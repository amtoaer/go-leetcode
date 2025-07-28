/*
 * @lc app=leetcode.cn id=708 lang=golang
 * @lcpr version=30202
 *
 * [708] 循环有序列表的插入
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Val  int
	Next *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		newNode := &Node{
			Val: x,
		}
		newNode.Next = newNode
		return newNode
	}
	current, lastNode := aNode, aNode
	for {
		if current.Val <= x && current.Next.Val >= x {
			current.Next = &Node{
				Val:  x,
				Next: current.Next,
			}
			break
		}
		if current.Val > current.Next.Val {
			lastNode = current
		}
		current = current.Next
		if current == aNode {
			lastNode.Next = &Node{
				Val:  x,
				Next: lastNode.Next,
			}
			break
		}
	}
	return aNode
}

// @lc code=end

func Test(t *testing.T) {
	createCircularList := func(vals []int) *Node {
		if len(vals) == 0 {
			return nil
		}
		head := &Node{Val: vals[0]}
		current := head
		for i := 1; i < len(vals); i++ {
			current.Next = &Node{Val: vals[i]}
			current = current.Next
		}
		current.Next = head
		return head
	}
	listToArray := func(head *Node) []int {
		if head == nil {
			return []int{}
		}
		result := []int{head.Val}
		current := head.Next
		for current != head {
			result = append(result, current.Val)
			current = current.Next
		}
		return result
	}

	tc := []struct {
		input []int
		x     int
		want  []int
	}{
		{[]int{3, 4, 1}, 2, []int{3, 4, 1, 2}},
		{[]int{}, 1, []int{1}},
		{[]int{1}, 0, []int{1, 0}},
	}

	for _, tt := range tc {
		head := createCircularList(tt.input)
		result := insert(head, tt.x)
		got := listToArray(result)
		assert.ElementsMatch(t, tt.want, got)
	}
}

/*
// @lcpr case=start
// [3,4,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// []\n1\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/
