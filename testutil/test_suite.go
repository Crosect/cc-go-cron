package testutil

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
	"reflect"
	"strings"
)

var jobMap map[string]cron.Job

func EnableCronTestSuite() fx.Option {
	return fx.Invoke(fx.Annotate(func(jobs []cron.Job) {
		jobMap = make(map[string]cron.Job, 0)
		for _, job := range jobs {
			jobName := strings.ToLower(getStructName(job))
			jobMap[jobName] = job
		}
	}, fx.ParamTags(`group:"cron_job"`)))
}

func RunJob(name string) {
	jobName := strings.ToLower(name)
	job, ok := jobMap[jobName]
	if !ok {
		panic(fmt.Errorf("job %v not found", jobName))
	}
	job.Run()
}

func getStructName(val interface{}) string {
	if t := reflect.TypeOf(val); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
