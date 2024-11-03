package main

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=638 lang=golang
 * @lcpr version=20002
 *
 * [638] 大礼包
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func shoppingOffers(price []int, special [][]int, needs []int) int {
	var filteredSpecial [][]int
	for _, spec := range special {
		var totalCnt, totalPrice int
		for idx, cnt := range spec[:len(spec)-1] {
			totalCnt += cnt
			totalPrice += cnt * price[idx]
		}
		if totalCnt != 0 && totalPrice > spec[len(spec)-1] {
			filteredSpecial = append(filteredSpecial, spec)
		}
	}
	byteNeeds := make([]byte, 0, len(needs))
	for _, need := range needs {
		byteNeeds = append(byteNeeds, byte(need))
	}
	var shopping func([]byte) int
	cache := make(map[string]int)
	shopping = func(needs []byte) int {
		if spent, ok := cache[string(needs)]; ok {
			return spent
		}
		// 初始化使用正常购买的价格
		var finalPrice int
		for idx, need := range needs {
			finalPrice += int(need) * price[idx]
		}
	label:
		for _, spec := range filteredSpecial {
			leftNeeds := make([]byte, 0, len(needs))
			for idx, cnt := range spec[:len(spec)-1] {
				left := int(needs[idx]) - cnt
				if left < 0 {
					continue label
				}
				leftNeeds = append(leftNeeds, byte(left))
			}
			finalPrice = min(finalPrice, spec[len(spec)-1]+shopping(leftNeeds))
		}
		cache[string(needs)] = finalPrice
		return finalPrice
	}
	return shopping(byteNeeds)
}

// @lc code=end

/*
// @lcpr case=start
// [2,5]\n[[3,0,5],[1,2,10]]\n[3,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n[[1,1,0,4],[2,2,1,9]]\n[1,2,1]\n
// @lcpr case=end

*/

func Test(t *testing.T) {
	tests := []struct {
		price   []int
		special [][]int
		needs   []int
		want    int
	}{
		{[]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}, 14},
		{[]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}, 11},
		{[]int{1, 1, 1}, [][]int{{1, 1, 1, 2}, {2, 2, 2, 3}}, []int{2, 2, 2}, 3},
		{[]int{2, 3}, [][]int{{1, 1, 4}, {2, 2, 7}}, []int{2, 2}, 7},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := shoppingOffers(tt.price, tt.special, tt.needs); got != tt.want {
				t.Errorf("shoppingOffers() = %v, want %v", got, tt.want)
			}
		})
	}
}
