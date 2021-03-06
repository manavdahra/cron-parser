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

// TODO: Decode documentation needed
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
