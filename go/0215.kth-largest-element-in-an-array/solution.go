// Created by amtoaer at 2025/11/21 10:51
// leetgo: 1.4.15
// https://leetcode.cn/problems/kth-largest-element-in-an-array/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	l, r, k := 0, n-1, n-k
	for {
		i := quickSelect(nums, l, r)
		if i == k {
			return nums[i]
		} else if i < k {
			l = i + 1
		} else {
			r = i - 1
		}
	}
}

func quickSelect(nums []int, l, r int) int {
	i := l + rand.Intn(r-l+1)
	pivot := nums[i]
	nums[i], nums[l] = nums[l], nums[i]
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
	return j
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := findKthLargest(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
