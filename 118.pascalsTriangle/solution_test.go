package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
 * @lc app=leetcode.cn id=118 lang=golang
 * @lcpr version=20004
 *
 * [118] 杨辉三角
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		if i == 0 {
			res = append(res, []int{1})
		} else if i == 1 {
			res = append(res, []int{1, 1})
		} else {
			beforeRow := res[len(res)-1]
			row := make([]int, 0, len(beforeRow)+1)
			row = append(row, 1)
			for j := 0; j < len(beforeRow)-1; j++ {
				row = append(row, beforeRow[j]+beforeRow[j+1])
			}
			row = append(row, 1)
			res = append(res, row)
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{5, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}},
		{1, [][]int{
			{1},
		}},
		{3, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
		}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("numRows=%d", tt.numRows), func(t *testing.T) {
			got := generate(tt.numRows)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate(%d) = %v, want %v", tt.numRows, got, tt.want)
			}
		})
	}
}
