package model

import (
	"deliveroo-cron/util"
	"fmt"
)

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
