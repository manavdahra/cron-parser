package model

import (
	"deliveroo-cron/common"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type Bitset struct {
	big.Int
	Min int
	Max int
}

// Decode is a function in Bitset struct
// This is used to decode a segment of a cron expression and set bits in big.Int field
//
// The result of this function is a bitset struct with big.Int field having bits set to 1 or 0
// 1 represent the numbers to be included in segment
// 0 represent the numbers to be excluded from segment
//
// For example:
// Given
// Days of week segment (Bitset)
// 0  0  0  0  0  0  0
// 0, 1, 2, 3, 4, 5, 6
//
// When
// Days of week = 1-5
//
// Then, after Decode
// 0  1  1  1  1  1  0
// 0, 1, 2, 3, 4, 5, 6
func (bitset *Bitset) Decode(expression string) error {
	fields := strings.Split(expression, ",")
	for _, field := range fields {
		if strings.Contains(field, "/") {
			split := strings.Split(field, "/")
			if len(split) > 2 {
				return common.ErrInvalidExpMultipleSteps
			}
			interval, err := NewInterval(split[0], bitset.Min, bitset.Max)
			if err != nil {
				return err
			}
			if !strings.Contains(split[0], "-") {
				interval.End = bitset.Max - 1
			}
			delta, err := strconv.Atoi(split[1])
			if err != nil {
				return err
			}
			if delta <= 0 {
				return common.ErrInvalidExpStepsNegative
			}
			for i := interval.Start; i <= interval.End; i += delta {
				bitset.SetBit(&bitset.Int, i, 1)
			}
		} else {
			interval, err := NewInterval(field, bitset.Min, bitset.Max)
			if err != nil {
				return err
			}
			for i := interval.Start; i <= interval.End; i++ {
				bitset.SetBit(&bitset.Int, i, 1)
			}
		}
	}
	return nil
}

func (bitset *Bitset) String() string {
	nums := make([]string, 0)
	for i := bitset.Min; i < bitset.Max; i++ {
		if bitset.Bit(i) != 0 {
			nums = append(nums, fmt.Sprintf("%d", i))
		}
	}
	return strings.Join(nums, ",")
}
