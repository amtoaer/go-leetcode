// Created by amtoaer at 2025/11/21 11:58
// leetgo: 1.4.15
// https://leetcode.cn/problems/sort-an-array/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func sortArray(nums []int) []int {
	l, r := 0, len(nums)-1
	quickSort(nums, l, r)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i := l + rand.Intn(r-l+1)
	pivot := nums[i]
	nums[l], nums[i] = nums[i], nums[l]
	i, j := l+1, r
	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[l], nums[j] = nums[j], nums[l]
	quickSort(nums, l, j-1)
	quickSort(nums, j+1, r)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := sortArray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
