// Created by amtoaer at 2025/11/23 18:43
// leetgo: 1.4.15
// https://leetcode.cn/problems/sliding-window-maximum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// 堆版本
// type Heap struct {
// 	nums  []int
// 	inner []int
// }

// func (h *Heap) Init() {
// 	for i := len(h.inner)/2 - 1; i >= 0; i-- {
// 		h.siftDown(i)
// 	}
// }

// func (h *Heap) siftDown(i int) {
// 	left, right := i*2+1, i*2+2
// 	biggest := i
// 	if left < len(h.inner) && h.nums[h.inner[left]] > h.nums[h.inner[biggest]] {
// 		biggest = left
// 	}
// 	if right < len(h.inner) && h.nums[h.inner[right]] > h.nums[h.inner[biggest]] {
// 		biggest = right
// 	}
// 	if biggest != i {
// 		h.inner[biggest], h.inner[i] = h.inner[i], h.inner[biggest]
// 		h.siftDown(biggest)
// 	}
// }

// func (h *Heap) siftUp(i int) {
// 	for i != 0 {
// 		parent := (i - 1) / 2
// 		if parent < 0 || h.nums[h.inner[parent]] > h.nums[h.inner[i]] {
// 			break
// 		}
// 		h.inner[parent], h.inner[i] = h.inner[i], h.inner[parent]
// 		i = parent
// 	}
// }

// func (h *Heap) Push(i int) {
// 	h.inner = append(h.inner, i)
// 	h.siftUp(len(h.inner) - 1)
// }

// func (h *Heap) Pop() int {
// 	res := h.inner[0]
// 	h.inner[0] = h.inner[len(h.inner)-1]
// 	h.inner = h.inner[:len(h.inner)-1]
// 	h.siftDown(0)
// 	return res
// }

// func (h *Heap) Top() int {
// 	return h.inner[0]
// }

// func maxSlidingWindow(nums []int, k int) (ans []int) {
// 	heapVal := make([]int, 0, k)
// 	for i := range k {
// 		heapVal = append(heapVal, i)
// 	}
// 	h := Heap{nums: nums, inner: heapVal}
// 	h.Init()
// 	ans = append(ans, nums[h.Top()])
// 	for i := k; i < len(nums); i++ {
// 		h.Push(i)
// 		for h.Top() <= i-k {
// 			h.Pop()
// 		}
// 		ans = append(ans, nums[h.Top()])
// 	}
// 	return
// }

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0, k)
	push := func(idx int) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[idx] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, idx)
	}
	for i := range k {
		push(i)
	}
	ans := make([]int, 0, len(nums)-k+1)
	ans = append(ans, nums[queue[0]])
	for i := k; i < len(nums); i++ {
		for len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		push(i)
		ans = append(ans, nums[queue[0]])
	}
	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := maxSlidingWindow(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
