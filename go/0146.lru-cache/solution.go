// Created by amtoaer at 2025/11/17 22:09
// leetgo: 1.4.15
// https://leetcode.cn/problems/lru-cache/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin
type LinkedListNode struct {
	key  int
	val  int
	prev *LinkedListNode
	next *LinkedListNode
}

type DoublyLinkedList struct {
	head, tail *LinkedListNode
}

func newDoublyLinkedList() *DoublyLinkedList {
	head, tail := &LinkedListNode{}, &LinkedListNode{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{
		head: head,
		tail: tail,
	}
}

func (d *DoublyLinkedList) remove(l *LinkedListNode) {
	prev, next := l.prev, l.next
	prev.next = next
	next.prev = prev
}

func (d *DoublyLinkedList) Insert(key, val int) *LinkedListNode {
	prev, next := d.head, d.head.next
	newNode := &LinkedListNode{key: key, val: val, prev: prev, next: next}
	prev.next, next.prev = newNode, newNode
	return newNode
}

func (d *DoublyLinkedList) MoveToHead(l *LinkedListNode) {
	d.remove(l)
	prev, next := d.head, d.head.next
	prev.next, l.prev = l, prev
	l.next, next.prev = next, l
}

func (d *DoublyLinkedList) RemoveLast() *LinkedListNode {
	last := d.tail.prev
	d.remove(last)
	return last
}

type LRUCache struct {
	m   map[int]*LinkedListNode
	dll *DoublyLinkedList
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:   make(map[int]*LinkedListNode),
		dll: newDoublyLinkedList(),
		cap: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.m[key]; ok {
		l.dll.MoveToHead(node)
		return node.val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.m[key]; ok {
		l.dll.MoveToHead(node)
		node.val = value
	} else {
		if len(l.m) == l.cap {
			last := l.dll.RemoveLast()
			delete(l.m, last.key)
		}
		newNode := l.dll.Insert(key, value)
		l.m[key] = newNode
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	constructorParams := MustSplitArray(params[0])
	capacity := Deserialize[int](constructorParams[0])
	obj := Constructor(capacity)

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "get":
			methodParams := MustSplitArray(params[i])
			key := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Get(key))
			output = append(output, ans)
		case "put":
			methodParams := MustSplitArray(params[i])
			key := Deserialize[int](methodParams[0])
			value := Deserialize[int](methodParams[1])
			obj.Put(key, value)
			output = append(output, "null")
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
