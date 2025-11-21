// Created by amtoaer at 2025/11/21 14:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/merge-k-sorted-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type Heap struct {
	inner []*ListNode
}

func (h *Heap) sinkDown(i int) {
	if i >= len(h.inner) {
		return
	}
	largest := i
	l, r := i*2+1, i*2+2
	if l < len(h.inner) && h.inner[l].Val < h.inner[largest].Val {
		largest = l
	}
	if r < len(h.inner) && h.inner[r].Val < h.inner[largest].Val {
		largest = r
	}
	if i != largest {
		h.inner[i], h.inner[largest] = h.inner[largest], h.inner[i]
		h.sinkDown(largest)
	}
}

func (h *Heap) Init() {
	for i := len(h.inner)/2 - 1; i >= 0; i-- {
		h.sinkDown(i)
	}
}

func (h *Heap) Pop() *ListNode {
	if len(h.inner) == 0 {
		return nil
	}
	top := h.inner[0]
	if top.Next != nil {
		h.inner[0] = top.Next
	} else {
		h.inner[0] = h.inner[len(h.inner)-1]
		h.inner = h.inner[:len(h.inner)-1]
	}
	h.sinkDown(0)
	return top
}

func mergeKLists(lists []*ListNode) *ListNode {
	innerHeap := make([]*ListNode, 0, len(lists))
	for _, node := range lists {
		if node == nil {
			continue
		}
		innerHeap = append(innerHeap, node)
	}
	heap := Heap{inner: innerHeap}
	heap.Init()
	dummy := &ListNode{}
	cur := dummy
	for {
		node := heap.Pop()
		if node == nil {
			return dummy.Next
		}
		cur.Next = node
		cur = cur.Next
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	lists := Deserialize[[]*ListNode](ReadLine(stdin))
	ans := mergeKLists(lists)

	fmt.Println("\noutput:", Serialize(ans))
}
