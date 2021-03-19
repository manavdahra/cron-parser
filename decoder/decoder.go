package decoder

import (
	"deliveroo-cron/common"
	"deliveroo-cron/model"
	"deliveroo-cron/util"
	"strings"
)

// decodeSequence is a list of all decoders to execute in a sequence
var decodeSequence = []func(field string, cron *model.Cron) error{
	decodeSeconds,
	decodeMinutes,
	decodeHours,
	decodeDaysOfMonth,
	decodeMonths,
	decodeDaysOfWeek,
}

func Decode(expression string, cron *model.Cron) error {
	if cron == nil {
		return common.ErrCronObjectNotDefined
	}
	cron.Expression = expression

	// split cron expression into segments
	segments := strings.Split(expression, " ")
	// ensure no. of segments = no. of decoders
	if len(segments) != len(decodeSequence) {
		return common.ErrInvalidExpMismatchedSeg
	}

	// decode all segments in a sequence. Fail fast if err.
	for index, field := range segments {
		if err := decodeSequence[index](field, cron); err != nil {
			return err
		}
	}

	return nil
}

func decodeSeconds(field string, cron *model.Cron) error {
	if err := cron.Seconds.Decode(field); err != nil {
		return err
	}
	return nil
}

func decodeMinutes(field string, cron *model.Cron) error {
	if err := cron.Minutes.Decode(field); err != nil {
		return err
	}
	return nil
}

func decodeHours(field string, cron *model.Cron) error {
	if err := cron.Hours.Decode(field); err != nil {
		return err
	}
	return nil
}

func decodeDaysOfMonth(field string, cron *model.Cron) error {
	if err := cron.DaysOfMonth.Decode(field); err != nil {
		return err
	}

	// clear 0th bit, as DaysOfMonth begins from 1
	cron.DaysOfMonth.SetBit(&cron.DaysOfMonth.Int, 0, 0)

	return nil
}

func decodeMonths(field string, cron *model.Cron) error {
	// ReplaceWithOrdinal replaces month name with an ordinal starting from 0. Example: APR -> 3
	if err := cron.Months.Decode(util.ReplaceWithOrdinal(field, common.Months)); err != nil {
		return err
	}

	// clear 0th bit, as Month begins from 1
	cron.Months.SetBit(&cron.Months.Int, 0, 0)

	return nil
}

func decodeDaysOfWeek(field string, cron *model.Cron) error {
	// ReplaceWithOrdinal replaces day name with an ordinal starting from 0. Example: THU -> 4
	if err := cron.DaysOfWeek.Decode(util.ReplaceWithOrdinal(field, common.DaysOfWeek)); err != nil {
		return err
	}
	return nil
}
