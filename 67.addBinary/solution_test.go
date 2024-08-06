package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	bytesA, bytesB := []byte(a), []byte(b)
	carry := 0
	bytes := []byte{}
	for carry > 0 || len(bytesA) > 0 || len(bytesB) > 0 {
		sum := carry
		if len(bytesA) > 0 {
			sum += int(bytesA[len(bytesA)-1] - '0')
			bytesA = bytesA[:len(bytesA)-1]

		}
		if len(bytesB) > 0 {
			sum += int(bytesB[len(bytesB)-1] - '0')
			bytesB = bytesB[:len(bytesB)-1]
		}
		carry = 0
		if sum >= 2 {
			sum -= 2
			carry = 1
		}
		bytes = append(bytes, '0'+byte(sum))
	}
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

// @lc code=end
func Test(t *testing.T) {
	tc := []struct {
		a, b string
		want string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{
			"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			"110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}
	for _, tt := range tc {
		assert.Equal(t, tt.want, addBinary(tt.a, tt.b))
	}
}
