package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

/*
 * @lc app=leetcode.cn id=690 lang=golang
 *
 * [690] Employee Importance
 */

// @lc code=start
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	m := make(map[int]*Employee)
	for _, e := range employees {
		m[e.Id] = e
	}
	queue := []int{id}
	var res int
	for len(queue) > 0 {
		lenQueue := len(queue)
		for _, id := range queue[:lenQueue] {
			e := m[id]
			res += e.Importance
			queue = append(queue, e.Subordinates...)
		}
		queue = queue[lenQueue:]
	}
	return res
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		input  []*Employee
		id     int
		output int
	}{
		{
			[]*Employee{
				{1, 5, []int{2, 3}},
				{2, 3, []int{}},
				{3, 3, []int{}},
			},
			1,
			11,
		},
		{
			[]*Employee{
				{1, 5, []int{2, 3}},
				{2, 3, []int{}},
				{3, 3, []int{}},
			},
			2,
			3,
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.output, getImportance(tt.input, tt.id))
	}
}
