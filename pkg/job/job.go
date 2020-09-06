package job

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// Status - Represents Status of a Job
type Status string

// list of status for a specific job
const (
	Paused      Status = "Paused"
	Running     Status = "Running"
	Initialized Status = "Queued"
	Done        Status = "Done"
	Terminated  Status = "Terminated"
)

// Job - Job represents a single entity of a job
type Job struct {
	sync.Mutex
	ID        string    `json:"ID"`
	Ts        time.Time `json:"Ts"`
	State     Status    `json:"State"`
	Data      []string  `json:"Data"`
	Err       error
	pause     chan struct{}
	resume    chan struct{}
	terminate chan struct{}
}

// MakeJob - Initiates the job
func MakeJob(data []string) Job {
	job := Job{
		ID:        uuid.New().String(),
		State:     Initialized,
		Ts:        time.Now(),
		Data:      data,
		pause:     make(chan struct{}),
		resume:    make(chan struct{}),
		terminate: make(chan struct{}),
	}
	return job
}
