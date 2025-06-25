/*
 * @lc app=leetcode.cn id=1089 lang=golang
 * @lcpr version=30201
 *
 * [1089] 复写零
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// @lc code=start
func duplicateZeros(arr []int) {
	var (
		j, top int
		i      = -1
	)
	for top < len(arr) {
		i++
		if arr[i] == 0 {
			top += 2
		} else {
			top++
		}
	}
	j = len(arr) - 1
	if top == len(arr)+1 {
		arr[j] = 0
		j--
		i--
	}
	for j >= 0 {
		arr[j] = arr[i]
		j--
		if arr[i] == 0 {
			arr[j] = arr[i]
			j--
		}
		i--
	}
}

// @lc code=end

func Test(t *testing.T) {
	tc := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tc {
		arrCopy := make([]int, len(tt.arr))
		copy(arrCopy, tt.arr)

		duplicateZeros(arrCopy)
		assert.Equal(t, tt.want, arrCopy)
	}
}

/*
// @lcpr case=start
// [1,0,2,3,0,4,5,0]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/
