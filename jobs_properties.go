package cccron

import (
	"github.com/crosect/cc-go/config"
)

type JobsProperties struct {
	Jobs map[string]JobProperties
}

func NewJobsProperties(loader config.Loader) (*JobsProperties, error) {
	props := &JobsProperties{}
	err := loader.Bind(props)
	return props, err
}

func (t *JobsProperties) Prefix() string {
	return "app.cron"
}
