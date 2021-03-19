package scheduler

import (
	"deliveroo-cron/model"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Duration time.Duration
	Period   int
}

type Scheduler struct {
	Crons    []model.Cron
	Config   Config
	Executor *Executor
	exit     chan os.Signal
}

func NewScheduler(config Config) *Scheduler {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	return &Scheduler{
		Config:   config,
		Executor: NewExecutor(),
		exit:     ch,
	}
}

func (scheduler *Scheduler) RegisterCron(cron *model.Cron) {
	if cron == nil {
		return
	}
	scheduler.Crons = append(scheduler.Crons, *cron)
}

func (scheduler *Scheduler) Run() {
	duration := scheduler.Config.Duration
	period := scheduler.Config.Period
	for {
		select {
		case code := <-scheduler.exit:
			switch code {
			case syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
				os.Exit(0)
			}
		default:
			now := time.Now()
			now = now.Truncate(time.Second)
			for _, cron := range scheduler.Crons {
				occurrence := cron.Next(now)
				if occurrence.Equal(now) {
					scheduler.Executor.Trigger <- &cron
				}
			}
			<-time.After(duration * time.Duration(period))
		}
	}
}
