# cc-go-cron

## Installation

```shell
go get github.com/crosect/cc-go-cron
```

## Usage

Create Job implement cron.Job interface

```go
package jobs

import (
	"github.com/robfig/cron/v3"
	"github.com/crosect/cc-go/log"
)

type FirstJob struct {
}

func NewFirstJob() cron.Job {
	return &FirstJob{}
}

// Run job handler
func (j *FirstJob) Run() {
	log.Infof("[FirstJob] job start")
	// TODO: implement logic here
	log.Infof("[FirstJob] job stop")
}
```

Provide job and provide cron to fx

```go
package bootstrap

import (
	"github.com/crosect/cc-go-cron"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(
		// Enable Cron
		cccron.EnableCron(),
		
		// Provide Jobs
		cccron.ProvideJob(jobs.NewFirstJob),
		cccron.ProvideJob(jobs.NewSecondJob),

		// Graceful shutdown.
		cccron.EnableGracefulShutdown(),
	)
}
```

Update YAML config

```yaml
app:
  cron:
    jobs:
      FirstJob:
        expression: "0 * * * * *" # Every minute at second 0
      SecondJob:
        expression: "@every 10s"
        disable: true
```

## Testing

Enable cc-go-cron Test Suite

```go
package testing

import (
	"github.com/crosect/cc-go-cron/testutil"
)

func Bootstrap() {
	fx.New(
		testutil.EnableCronTestSuite(),
	)
}
```

Run a job

```go
package testing

import (
	"github.com/crosect/cc-go-cron/testutil"
)

func RunJob() {
	testutil.RunJob("JobName")
}
```
