// Created by amtoaer at 2025/12/17 23:52
// leetgo: 1.4.15
// https://leetcode.cn/problems/reverse-integer/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverse(x int) (ans int) {
	shouldReverse := false
	if x < 0 {
		x = -x
		shouldReverse = true
	}
	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if shouldReverse {
		ans = -ans
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := reverse(x)

	fmt.Println("\noutput:", Serialize(ans))
}
