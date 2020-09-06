package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/someshkoli/jobHandler/pkg/job"
)

// JobHandler - pipeline job handler
type JobHandler struct {
	*sync.Mutex
	JobPool job.Pool `json:"JobPool"`
}

// MakeJobHandler - returns instance of a new job handler
func MakeJobHandler() JobHandler {
	pool := make(job.Pool)
	return JobHandler{
		JobPool: pool,
	}
}

// RegisterHandlers - registers routes and respective handler and returns mux for the same
func (JH *JobHandler) RegisterHandlers() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", JH.ping)
	mux.HandleFunc("/upload", JH.upload)
	mux.HandleFunc("/pause", JH.pause)
	mux.HandleFunc("/resume", JH.resume)
	mux.HandleFunc("/terminate", JH.terminate)
	mux.HandleFunc("/status", JH.getStatus)
	return mux
}

// StartServer - set server for the job handler
func (JH *JobHandler) StartServer(srv *http.Server) error {
	fmt.Println("Starting Server")
	return srv.ListenAndServe()
}
