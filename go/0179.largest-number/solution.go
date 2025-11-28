// Created by amtoaer at 2025/11/29 01:09
// leetgo: 1.4.15
// https://leetcode.cn/problems/largest-number/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		numI, numJ := nums[i], nums[j]
		strI := strconv.Itoa(numI)
		strJ := strconv.Itoa(numJ)
		return strI+strJ > strJ+strI
	})
	var sb strings.Builder
	for _, num := range nums {
		sb.WriteString(strconv.Itoa(num))
	}
	res := sb.String()
	if num, err := strconv.Atoi(res); err == nil && num == 0 {
		return "0"
	}
	return sb.String()
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := largestNumber(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
