// Created by amtoaer at 2025/11/21 12:08
// leetgo: 1.4.15
// https://leetcode.cn/problems/add-strings/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func addStrings(num1 string, num2 string) string {
	nums1Slice, nums2Slice := []byte(num1), []byte(num2)
	reverse(nums1Slice)
	reverse(nums2Slice)
	if len(nums1Slice) < len(nums2Slice) {
		nums1Slice, nums2Slice = nums2Slice, nums1Slice
	}
	// 确保 nums1Slice 长度大于等于 nums2Slice
	var idx, carry int
	lenNums1 := len(nums1Slice)
	for idx < lenNums1 || carry != 0 {
		sum := carry
		if idx < len(nums1Slice) {
			sum += int(nums1Slice[idx] - '0')
		}
		if idx < len(nums2Slice) {
			sum += int(nums2Slice[idx] - '0')
		}
		curByte := byte(sum%10) + '0'
		if idx < lenNums1 {
			nums1Slice[idx] = curByte
		} else {
			nums1Slice = append(nums1Slice, curByte)
		}
		carry = sum / 10
		idx++
	}
	reverse(nums1Slice)
	return string(nums1Slice)
}

func reverse(nums []byte) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	num1 := Deserialize[string](ReadLine(stdin))
	num2 := Deserialize[string](ReadLine(stdin))
	ans := addStrings(num1, num2)

	fmt.Println("\noutput:", Serialize(ans))
}
