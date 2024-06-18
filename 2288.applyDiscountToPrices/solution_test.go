package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=2288 lang=golang
 *
 * [2288] Apply Discount to Prices
 */

// @lc code=start

func discountItem(sentence string, discount int) string {
	if len(sentence) == 0 || sentence[0] != '$' {
		return sentence
	}
	if len(sentence) == 1 {
		return sentence
	}
	var res int
	for i := 1; i < len(sentence); i++ {
		if sentence[i] > '9' || sentence[i] < '0' {
			return sentence
		}
		res = res*10 + int(sentence[i]-'0')
	}
	return fmt.Sprintf("$%.2f", float64(res*(100-discount))*0.01)
}

func discountPrices(sentence string, discount int) string {
	var sb strings.Builder
	start, end := 0, 0
	for end <= len(sentence) {
		if end == len(sentence) || sentence[end] == ' ' {
			i := discountItem(sentence[start:end], discount)
			sb.WriteString(i)
			if end != len(sentence) {
				sb.WriteByte(' ')
			}
			start = end + 1
		}
		end++
	}
	return sb.String()
}

// @lc code=end

func Test(t *testing.T) {
	testcases := []struct {
		sentence string
		discount int
		want     string
	}{
		{sentence: "The price is $100", discount: 20, want: "The price is $80.00"},
		{sentence: "The price is $100", discount: 10, want: "The price is $90.00"},
		{
			sentence: "1 2 $3 4 $5 $6 7 8$ $9 $10$",
			discount: 100,
			want:     "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$",
		},
		{
			sentence: "there are $1 $2 and 5$ candies in the shop",
			discount: 50,
			want:     "there are $0.50 $1.00 and 5$ candies in the shop",
		},
		{
			sentence: "$7383692 5q $5870426",
			discount: 64,
			want:     "$2658129.12 5q $2113353.36",
		},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.want, discountPrices(tc.sentence, tc.discount))
	}
}
