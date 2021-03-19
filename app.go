package main

import (
	"deliveroo-cron/common"
	"deliveroo-cron/decoder"
	"deliveroo-cron/model"
	"deliveroo-cron/scheduler"
	"deliveroo-cron/util"
	"fmt"
	"os"
	"time"
)

func main() {
	appArguments := os.Args

	// Validate, we have right no of program arguments
	if len(appArguments) != 2 {
		panic(fmt.Sprintf(common.HelpFmt, common.CronExp))
	}

	cronExpression, command := util.ParseCommandLineArgs(appArguments[1])

	// instantiate a new empty Trigger struct
	cron := model.NewCron(command)

	// decode it using cron expression
	err := decoder.Decode(cronExpression, cron)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(cron.String())
	fmt.Println(cron.Next(time.Now()))
	_scheduler := scheduler.NewScheduler(scheduler.Config{
		Duration: time.Millisecond,
		Period:   500,
	})
	_scheduler.RegisterCron(cron)
	_scheduler.Run()
}
