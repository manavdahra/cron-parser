package scheduler

import (
	"deliveroo-cron/model"
	"fmt"
	"time"
)

type Executor struct {
	Trigger chan *model.Cron
}

func NewExecutor() *Executor {
	exec := &Executor{Trigger: make(chan *model.Cron)}
	exec.Subscribe()
	return exec
}

func (task *Executor) Subscribe() {
	go func() {
		for {
			select {
			case cron := <-task.Trigger:
				if cron != nil {
					task.execute(*cron)
				}
			}
		}
	}()
}

func (task *Executor) execute(cron model.Cron) {
	go func() {
		fmt.Println(fmt.Sprintf("starting cron job: %s @ %s", cron.Command, time.Now()))
		// TODO: execute the job
		fmt.Println(fmt.Sprintf("finished cron job: %s @ %s", cron.Command, time.Now()))
	}()
}
