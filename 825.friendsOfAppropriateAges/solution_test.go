package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=825 lang=golang
 * @lcpr version=20003
 *
 * [825] 适龄的朋友
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numFriendRequests(ages []int) int {
	var nums, prev [121]int
	for _, age := range ages {
		nums[age]++
	}
	for i := 1; i <= 120; i++ {
		prev[i] = prev[i-1] + nums[i]
	}
	res := 0
	for i := 15; i <= 120; i++ {
		if nums[i] == 0 {
			continue
		}
		left := i/2 + 7
		res += nums[i] * (prev[i] - prev[left] - 1)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [16,16]\n
// @lcpr case=end

// @lcpr case=start
// [16,17,18]\n
// @lcpr case=end

// @lcpr case=start
// [20,30,100,110,120]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	type Case struct {
		Input  []int
		Expect int
	}
	cases := []Case{
		{Input: []int{16, 16}, Expect: 2},
		{Input: []int{16, 17, 18}, Expect: 2},
		{Input: []int{20, 30, 100, 110, 120}, Expect: 3},
	}
	for i, c := range cases {
		output := numFriendRequests(c.Input)
		if output != c.Expect {
			t.Fatalf("case %d fail expect %v got %v", i, c.Expect, output)
		}
	}
}
