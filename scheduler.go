package cccron

import (
	"context"
	"fmt"
	"github.com/crosect/cc-go/web/log"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
	"strings"
)

type Scheduler struct {
	c *cron.Cron
}

type SchedulerConstructorParams struct {
	fx.In
	Properties *JobsProperties
	Jobs       []cron.Job `group:"cron_job"`
}

func NewScheduler(params SchedulerConstructorParams) (*Scheduler, error) {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(cron.DefaultLogger),
			cron.SkipIfStillRunning(cron.DefaultLogger),
		),
	)
	s := &Scheduler{c: c}
	for _, job := range params.Jobs {
		jobName := getStructName(job)
		if jobProperties, ok := params.Properties.Jobs[strings.ToLower(jobName)]; ok {
			if !jobProperties.Disable {
				_, err := c.AddJob(jobProperties.Expression, job)
				if err != nil {
					return nil, fmt.Errorf("[cccron] could not add %s to scheduler: %s", jobName, err)
				}
				log.Infof("[cccron] Add job %s to scheduler successfully.", jobName)
			}
		}
	}
	return s, nil
}

func (s *Scheduler) Start() {
	s.c.Start()
}

func (s *Scheduler) Stop() error {
	ctx := s.c.Stop()
	for {
		if ctx.Err() == context.Canceled {
			return nil
		}
	}
}
