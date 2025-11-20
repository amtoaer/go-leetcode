// Created by amtoaer at 2025/11/18 14:53
// leetgo: 1.4.15
// https://leetcode.cn/problems/3sum/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	next := func(idx int) int {
		next := idx
		for next < len(nums) && nums[next] == nums[idx] {
			next++
		}
		return next
	}
	prev := func(idx int) int {
		prev := idx
		for prev >= 0 && nums[prev] == nums[idx] {
			prev--
		}
		return prev
	}
	for i := 0; i < len(nums); i = next(i) {
		j, k := i+1, len(nums)-1
		target := -1 * nums[i]
		for j < k {
			cur := nums[j] + nums[k]
			if cur > target {
				k = prev(k)
			} else if cur < target {
				j = next(j)
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j = next(j)
			}
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := threeSum(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
