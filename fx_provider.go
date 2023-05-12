package cccron

import (
	"context"
	"github.com/crosect/cc-go"
	"github.com/crosect/cc-go/log"
	"go.uber.org/fx"
)

func ProvideJob(job interface{}) fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  "cron_job",
		Target: job,
	})
}

func EnableCron() fx.Option {
	return fx.Options(
		ccgo.ProvideProps(NewJobsProperties),
		fx.Provide(NewScheduler),
		fx.Invoke(func(lc fx.Lifecycle, scheduler *Scheduler) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					scheduler.Start()
					return nil
				},
			})
		}),
	)
}

func EnableGracefulShutdown() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, scheduler *Scheduler) {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				log.Info("[cccron] Stopping Scheduler")
				if err := scheduler.Stop(); err != nil {
					log.Errorf("[cccron] Could not stop cron: %v", err)
					return err
				}
				log.Info("[cccron] Scheduler Stopped")
				return nil
			},
		})
	})
}
