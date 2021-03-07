package decoder

import (
	"deliveroo-cron/common"
	"deliveroo-cron/model"
	"deliveroo-cron/util"
	"strings"
)

var decodeSequence = []func(field string, cron *model.Cron) error{
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
	fields := strings.Split(expression, " ")
	if len(fields) != len(decodeSequence) {
		return common.ErrInvalidExpMismatchedSeg
	}

	for index, field := range fields {
		if err := decodeSequence[index](field, cron); err != nil {
			return err
		}
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
	cron.DaysOfMonth.SetBit(&cron.DaysOfMonth.Int, 0, 0)
	return nil
}

func decodeMonths(field string, cron *model.Cron) error {
	if err := cron.Months.Decode(util.ReplaceWithOrdinal(field, common.Months)); err != nil {
		return err
	}
	cron.Months.SetBit(&cron.Months.Int, 0, 0)
	return nil
}

func decodeDaysOfWeek(field string, cron *model.Cron) error {
	if err := cron.DaysOfWeek.Decode(util.ReplaceWithOrdinal(field, common.DaysOfWeek)); err != nil {
		return err
	}
	return nil
}
