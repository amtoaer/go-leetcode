package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] Find Median from Data Stream
 */

// @lc code=start

func bigger(i, j int) bool {
	return i > j
}

func less(i, j int) bool {
	return i < j
}

func sinkDown(arr []int, start, end int, better func(i, j int) bool) {
	cur, child := start, start*2+1
	for child <= end {
		if child+1 <= end && better(arr[child+1], arr[child]) {
			child++
		}
		if better(arr[cur], arr[child]) {
			return
		}
		arr[cur], arr[child] = arr[child], arr[cur]
		cur, child = child, child*2+1
	}
}

func sinkUp(arr []int, end int, better func(i, j int) bool) {
	cur, parent := end, (end-1)/2
	for cur > 0 && better(arr[cur], arr[parent]) {
		arr[cur], arr[parent] = arr[parent], arr[cur]
		cur, parent = parent, (parent-1)/2
	}
}

type MedianFinder struct {
	heapMin []int // 存放小的一半的大顶堆
	heapMax []int // 存放大的一半的小顶堆
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (m *MedianFinder) AddNum(num int) {
	if len(m.heapMin) == 0 || num <= m.heapMin[0] {
		// 没有元素，或者数字小于“小的一半的最大值”
		m.heapMin = append(m.heapMin, num)
		sinkUp(m.heapMin, len(m.heapMin)-1, bigger)
		if len(m.heapMax)+1 < len(m.heapMin) {
			// 把小的部分中的最大值放到大的一半里
			v := m.heapMin[0]
			m.heapMin[0], m.heapMin[len(m.heapMin)-1] = m.heapMin[len(m.heapMin)-1], m.heapMin[0]
			m.heapMin = m.heapMin[:len(m.heapMin)-1]
			sinkDown(m.heapMin, 0, len(m.heapMin)-1, bigger)

			m.heapMax = append(m.heapMax, v)
			sinkUp(m.heapMax, len(m.heapMax)-1, less)
		}
	} else {
		// 数字大于小的一半的最大值，应该放到大的一半里
		m.heapMax = append(m.heapMax, num)
		sinkUp(m.heapMax, len(m.heapMax)-1, less)

		if len(m.heapMax) > len(m.heapMin) {
			v := m.heapMax[0]
			m.heapMax[0], m.heapMax[len(m.heapMax)-1] = m.heapMax[len(m.heapMax)-1], m.heapMax[0]
			m.heapMax = m.heapMax[:len(m.heapMax)-1]
			sinkDown(m.heapMax, 0, len(m.heapMax)-1, less)

			m.heapMin = append(m.heapMin, v)
			sinkUp(m.heapMin, len(m.heapMin)-1, bigger)
		}
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if len(m.heapMin) > len(m.heapMax) {
		return float64(m.heapMin[0])
	}
	return float64(m.heapMin[0]+m.heapMax[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

func Test(t *testing.T) {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	if obj.FindMedian() != 1.5 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(3)
	if obj.FindMedian() != 2 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(4)
	if obj.FindMedian() != 2.5 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 2 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 1.5 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 1 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0.5 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0 {
		t.Fatal(obj.FindMedian())
	}
	obj.AddNum(0)
	if obj.FindMedian() != 0 {
		t.Fatal(obj.FindMedian())
	}
}
