package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2806 lang=golang
 *
 * [2806] Account Balance After Rounded Purchase
 */

// @lc code=start
func accountBalanceAfterPurchase(purchaseAmount int) int {
	x := purchaseAmount % 10
	var money int
	if x >= 5 {
		money = purchaseAmount + 10 - x
	} else {
		money = purchaseAmount - x
	}
	return 100 - money
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		purchaseAmount int
		expected       int
	}{
		{purchaseAmount: 9, expected: 90},
		{purchaseAmount: 100, expected: 0},
		{purchaseAmount: 15, expected: 80},
	}
	for _, testcase := range testcases {
		assert.Equal(t, testcase.expected, accountBalanceAfterPurchase(testcase.purchaseAmount))
	}
}
