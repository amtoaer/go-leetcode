// Created by amtoaer at 2025/11/21 13:00
// leetgo: 1.4.15
// https://leetcode.cn/problems/merge-intervals/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func merge(intervals [][]int) (ans [][]int) {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		nextLeft, nextRight := intervals[i][0], intervals[i][1]
		if nextRight > right && nextLeft <= right {
			right = nextRight
		} else if nextLeft > right {
			ans = append(ans, []int{left, right})
			left, right = nextLeft, nextRight
		}
		if i == len(intervals)-1 {
			ans = append(ans, []int{left, right})
		}
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	intervals := Deserialize[[][]int](ReadLine(stdin))
	ans := merge(intervals)

	fmt.Println("\noutput:", Serialize(ans))
}
