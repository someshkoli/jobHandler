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
	ID        string   `json:"ID"`
	Ts        string   `json:"Ts"`
	State     Status   `json:"State"`
	Data      []string `json:"Data"`
	Err       error
	pause     chan struct{}
	resume    chan struct{}
	terminate chan struct{}
}

// MakeJob - Initiates the job
func MakeJob(data []string) Job {
	id := uuid.New().String()
	time := time.Now().Format("2 Jan 2006 15:04:05")
	job := Job{
		ID:        id,
		State:     Initialized,
		Ts:        time,
		Data:      data,
		pause:     make(chan struct{}),
		resume:    make(chan struct{}),
		terminate: make(chan struct{}),
	}
	return job
}
