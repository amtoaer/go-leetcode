// Created by amtoaer at 2025/11/22 10:59
// leetgo: 1.4.15
// https://leetcode.cn/problems/compare-version-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func compareVersion(version1 string, version2 string) int {
	version1Bytes, version2Bytes := []byte(version1), []byte(version2)
	var idx1, idx2 int
	for idx1 < len(version1Bytes) || idx2 < len(version2Bytes) {
		var v1, v2 int
		for idx1 < len(version1Bytes) && version1Bytes[idx1] != '.' {
			v1 = v1*10 + int(version1Bytes[idx1]-'0')
			idx1++
		}
		for idx2 < len(version2Bytes) && version2Bytes[idx2] != '.' {
			v2 = v2*10 + int(version2Bytes[idx2]-'0')
			idx2++
		}
		if v1 < v2 {
			return -1
		}
		if v1 > v2 {
			return 1
		}
		// 跳过 .
		idx1++
		idx2++
	}
	return 0
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	version1 := Deserialize[string](ReadLine(stdin))
	version2 := Deserialize[string](ReadLine(stdin))
	ans := compareVersion(version1, version2)

	fmt.Println("\noutput:", Serialize(ans))
}
