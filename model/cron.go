package model

import (
	"deliveroo-cron/util"
	"fmt"
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
	}
}

func (cron *Cron) String() string {
	return fmt.Sprintf(util.GenerateOutputFmt(), cron.Minutes.String(),
		cron.Hours.String(),
		cron.DaysOfMonth.String(),
		cron.Months.String(),
		cron.DaysOfWeek.String(),
		cron.Command)
}
