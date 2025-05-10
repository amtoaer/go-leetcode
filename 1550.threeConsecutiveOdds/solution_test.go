/*
 * @lc app=leetcode.cn id=1550 lang=golang
 * @lcpr version=30201
 *
 * [1550] 存在连续三个奇数的数组
 */

package main

import (
	"testing"
)

// @lc code=start
func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	var cnt, left int
	for right, rightNum := range arr {
		if rightNum&1 == 1 {
			cnt++
		}
		for right-left >= 3 {
			if arr[left]&1 == 1 {
				cnt--
			}
			left++
		}
		if cnt == 3 {
			return true
		}
	}
	return false
}

// @lc code=end

func Test(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{2, 6, 4, 1}, false},
		{[]int{1, 2, 34, 3, 4, 5, 7, 23, 12}, true},
		{[]int{1, 3, 5}, true},
		{[]int{1, 2, 1, 1}, false},
		{[]int{1, 3, 4}, false},
		{[]int{}, false},
		{[]int{1, 2}, false},
		{[]int{2, 4, 6}, false},
		{[]int{1, 3, 5, 7}, true},
		{[]int{2, 1, 3, 5}, true},
		{[]int{2, 1, 3, 4}, false},
	}
	for _, tt := range tests {
		if got := threeConsecutiveOdds(tt.arr); got != tt.want {
			t.Errorf("threeConsecutiveOdds(%v) = %v, want %v", tt.arr, got, tt.want)
		}
	}
}

/*
// @lcpr case=start
// [2,6,4,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,34,3,4,5,7,23,12]\n
// @lcpr case=end

*/
