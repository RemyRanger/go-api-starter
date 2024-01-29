package entities

import "time"

// Job defines model for Job.
type Job struct {
	ExecDate time.Time
	Context  string
	Name     string
	Status   string
	Duration int64
}
