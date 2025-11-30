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
	if rand.Intn(2) == 1 {
		quickSort(nums)
	} else {
		heapSort(nums)
	}
	return nums
}

// / 快速排序
func quickSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	i := rand.Intn(len(nums))
	pivot := nums[i]
	nums[0], nums[i] = nums[i], nums[0]
	i, j := 1, len(nums)-1
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
	nums[0], nums[j] = nums[j], nums[0]
	quickSort(nums[:j])
	quickSort(nums[j+1:])
}

// / 堆排序
func heapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, i, len(nums))
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		siftDown(nums, 0, i)
	}
}

func siftDown(nums []int, i, length int) {
	for i < length {
		biggest := i
		left, right := i*2+1, i*2+2
		if left < length && nums[left] > nums[biggest] {
			biggest = left
		}
		if right < length && nums[right] > nums[biggest] {
			biggest = right
		}
		if biggest == i {
			break
		}
		nums[i], nums[biggest] = nums[biggest], nums[i]
		i = biggest
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := sortArray(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
