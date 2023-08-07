package cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"rk-grpc/internal/global"
	_ "rk-grpc/internal/infra/task"
)

var Cron *cron.Cron

func InitTask() {
	c := cron.New()

	for s, t := range global.Tasks {
		fmt.Printf("add task: %s", s)
		c.AddFunc("@every 1s", t.Execute)
	}

	if len(c.Entries()) == 0 {
		return
	}
	c.Start()
}
