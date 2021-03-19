package model

import (
	"deliveroo-cron/util"
	"fmt"
	"time"
)

// Cron struct is used to represent the cron expression in expanded form
// For example:
// Given,
// cron expression: */15 0 1,15 * 1-5
//
// that is,
// Minutes segment: 		*/15
// Hours segment: 			0
// DaysOfMonth segment: 	1,15
// Months segment: 			*
// DaysOfWeek segment: 		1-5
//
// Then,
// Bitset is used to represent each segment field in Cron in expanded form
// it has a big.Int field having bits set to either 1 or 0
// 1 represent the numbers to be included in segment
// 0 represent the numbers to be excluded from segment
// For example:
// DaysOfWeek = 1-5
//
// Then,
// DaysOfWeek Bitset
// 0  1  1  1  1  1  0
// 0, 1, 2, 3, 4, 5, 6
type Cron struct {
	Expression  string
	Months      Bitset
	DaysOfMonth Bitset
	DaysOfWeek  Bitset
	Hours       Bitset
	Minutes     Bitset
	Seconds     Bitset
	Command     string
}

func NewCron(command string) *Cron {
	return &Cron{
		Command:     command,
		Months:      Bitset{Min: 0, Max: 13},
		DaysOfMonth: Bitset{Min: 0, Max: 32},
		DaysOfWeek:  Bitset{Min: 0, Max: 7},
		Hours:       Bitset{Min: 0, Max: 24},
		Minutes:     Bitset{Min: 0, Max: 60},
		Seconds:     Bitset{Min: 0, Max: 60},
	}
}

func (cron *Cron) Next(occurrence time.Time) time.Time {
	occurrence = occurrence.Truncate(time.Second)
	nextSec := cron.Seconds.NextSetBit(occurrence.Second())
	if nextSec != occurrence.Second() {
		minute := occurrence.Minute()
		if nextSec == -1 {
			minute++
			nextSec = 0
		}
		occurrence = time.Date(occurrence.Year(), occurrence.Month(), occurrence.Day(), occurrence.Hour(), minute, nextSec, 0, occurrence.Location())
		occurrence = cron.Next(occurrence)
	} else {
		occurrence = time.Date(occurrence.Year(), occurrence.Month(), occurrence.Day(), occurrence.Hour(), occurrence.Minute(), nextSec, 0, occurrence.Location())
	}

	nextMin := cron.Minutes.NextSetBit(occurrence.Minute())
	if nextMin != occurrence.Minute() {
		hour := occurrence.Hour()
		if nextMin == -1 {
			hour++
			nextMin = 0
		}
		occurrence = time.Date(occurrence.Year(), occurrence.Month(), occurrence.Day(), hour, nextMin, 0, 0, occurrence.Location())
		occurrence = cron.Next(occurrence)
	} else {
		occurrence = time.Date(occurrence.Year(), occurrence.Month(), occurrence.Day(), occurrence.Hour(), nextMin, occurrence.Second(), 0, occurrence.Location())
	}

	nextHr := cron.Hours.NextSetBit(occurrence.Hour())
	if nextHr != occurrence.Hour() {
		day := occurrence.Day()
		if nextHr == -1 {
			day++
			nextHr = 0
		}
		occurrence = time.Date(occurrence.Year(), occurrence.Month(), day, nextHr, 0, 0, 0, occurrence.Location())
		occurrence = cron.Next(occurrence)
	} else {
		occurrence = time.Date(occurrence.Year(), occurrence.Month(), occurrence.Day(), nextHr, occurrence.Minute(), occurrence.Second(), 0, occurrence.Location())
	}

	nextDayOfMonth := cron.DaysOfMonth.NextSetBit(occurrence.Day())
	nextDayOfWeek := cron.DaysOfWeek.NextSetBit(int(occurrence.Weekday()))
	if nextDayOfMonth != occurrence.Day() {
		month := occurrence.Month()
		if nextDayOfMonth == -1 || nextDayOfWeek == -1 {
			month++
			nextDayOfMonth = 1
		}
		occurrence = time.Date(occurrence.Year(), month, nextDayOfMonth, 0, 0, 0, 0, occurrence.Location())
		occurrence = cron.Next(occurrence)
	} else {
		occurrence = time.Date(occurrence.Year(), occurrence.Month(), nextDayOfMonth, occurrence.Hour(), occurrence.Minute(), occurrence.Second(), 0, occurrence.Location())
	}

	nextMonth := cron.Months.NextSetBit(int(occurrence.Month()))
	if nextMonth != int(occurrence.Month()) {
		year := occurrence.Year()
		if nextMonth == -1 {
			year++
			nextMonth = 1
		}
		occurrence = time.Date(year, time.Month(nextMonth), 1, 0, 0, 0, 0, occurrence.Location())
		occurrence = cron.Next(occurrence)
	} else {
		occurrence = time.Date(occurrence.Year(), time.Month(nextMonth), occurrence.Day(), occurrence.Hour(), occurrence.Minute(), occurrence.Second(), 0, occurrence.Location())
	}
	return occurrence
}

func (cron *Cron) String() string {
	return fmt.Sprintf(util.GenerateOutputFmt(), cron.Seconds.String(),
		cron.Minutes.String(),
		cron.Hours.String(),
		cron.DaysOfMonth.String(),
		cron.Months.String(),
		cron.DaysOfWeek.String(),
		cron.Command)
}
