package model

import (
	"deliveroo-cron/common"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

// TODO: NewInterval documentation needed
func NewInterval(field string, min, max int) (*Interval, error) {
	interval := &Interval{}
	if strings.Contains(field, "*") {
		interval.Start = min
		interval.End = max - 1
		return interval, nil
	}
	if strings.Contains(field, "-") {
		split := strings.Split(field, "-")
		if len(split) > 2 {
			return nil, common.ErrInvalidIntervalRanges
		}
		beg, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		interval.Start = beg
		interval.End = end
	} else {
		val, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		interval.Start = val
		interval.End = val
	}
	if interval.Start < min || interval.End < min || interval.Start >= max || interval.End >= max {
		return nil, common.ErrIntervalOutOfBounds
	}
	if interval.Start > interval.End {
		return nil, common.ErrInvalidIntervalLimits
	}
	return interval, nil
}