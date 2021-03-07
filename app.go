package main

import (
	"deliveroo-cron/common"
	"deliveroo-cron/decoder"
	"deliveroo-cron/model"
	"deliveroo-cron/util"
	"fmt"
	"os"
)

func main() {
	appArguments := os.Args

	helperString := fmt.Sprintf(common.HelpFmt, common.CronExp)
	if len(appArguments) != 2 {
		panic(helperString)
	}

	cronExpression, command := util.ParseCommandLineArgs(appArguments[1])

	cron := model.NewCron(command)
	err := decoder.Decode(cronExpression, cron)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(cron.String())
}
