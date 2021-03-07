package model

import (
	"deliveroo-cron/common"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInterval_Valid_Simple(t *testing.T) {
	day := 15
	min := 1
	max := 32

	interval, err := NewInterval(fmt.Sprintf("%d", day), min, max)
	assert.NoError(t, err, nil)
	assert.Equal(t, day, interval.Start)
	assert.Equal(t, day, interval.End)
}

func TestNewInterval_Valid_All(t *testing.T) {
	min := 1
	max := 32

	interval, err := NewInterval("*", min, max)
	assert.NoError(t, err, nil)
	assert.Equal(t, min, interval.Start)
	assert.Equal(t, max-1, interval.End)
}

func TestNewInterval_Valid_Range(t *testing.T) {
	start := 1
	end := 30

	min := 1
	max := 32

	interval, err := NewInterval(fmt.Sprintf("%d-%d", start, end), min, max)
	assert.NoError(t, err, nil)
	assert.Equal(t, start, interval.Start)
	assert.Equal(t, end, interval.End)
}

func TestNewInterval_Invalid_OutOfBounds(t *testing.T) {
	start := 1
	end := 33

	min := 1
	max := 32

	interval, err := NewInterval(fmt.Sprintf("%d-%d", start, end), min, max)
	assert.Error(t, err)
	assert.Nil(t, interval)
	assert.Equal(t, common.ErrIntervalOutOfBounds, err)

	start = 0
	end = 31

	interval, err = NewInterval(fmt.Sprintf("%d-%d", start, end), min, max)
	assert.Error(t, err)
	assert.Nil(t, interval)
	assert.Equal(t, common.ErrIntervalOutOfBounds, err)
}

func TestNewInterval_Invalid_Range(t *testing.T) {
	start := 1
	end1 := 15
	end2 := 31

	min := 1
	max := 32

	interval, err := NewInterval(fmt.Sprintf("%d-%d-%d", start, end1, end2), min, max)
	assert.Error(t, err)
	assert.Nil(t, interval)
	assert.Equal(t, common.ErrInvalidIntervalRanges, err)
}
