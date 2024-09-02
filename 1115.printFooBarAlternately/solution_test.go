package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=1115 lang=golang
 *
 * [1115] Print FooBar Alternately
 */

// @lc code=start

func printFoo() {
	print("foo")
}

func printBar() {
	print("bar")
}

type FooBar struct {
	n   int
	foo chan struct{}
	bar chan struct{}
}

func NewFooBar(n int) *FooBar {
	foobar := &FooBar{
		n, make(chan struct{}, 1), make(chan struct{}),
	}
	foobar.foo <- struct{}{}
	return foobar
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.foo
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.bar <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.bar
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.foo <- struct{}{}
	}
}

// @lc code=end

func Test(t *testing.T) {
	fb := NewFooBar(5)
	go fb.Foo(printFoo)
	go fb.Bar(printBar)
}
