package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/someshkoli/jobHandler/pkg/job"
)

// Ping - Handler for ping route
// Pings the server
func (JH *JobHandler) ping(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	respond(res, "Pong", true, http.StatusOK)
}

// upload - Handler for /upload
// starts upload job
func (JH *JobHandler) upload(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	var buf bytes.Buffer

	req.ParseMultipartForm(100)

	file, handler, err := req.FormFile("file")

	if err != nil {
		respond(res, "Could not read input file", true, http.StatusBadRequest)
	}
	count, err := io.Copy(&buf, file)
	if err != nil {
		respond(res, "Could not read input file", true, http.StatusInternalServerError)
	}

	fmt.Printf("[UPLOADING-%d-Characters][%s][%s]\n", count, handler.Filename, time.Now().Format("2 Jan 2006 15:04:05"))
	contents := buf.String()

	newJob := job.MakeJob(strings.Split(contents, "\n"))
	currentJob := JH.JobPool.Add(newJob.ID, &newJob)
	currentJob.Start()
	respond(res, currentJob.ID, true, http.StatusOK)
}

// pause - Handler for /pause
// pauses a job with given job id
func (JH *JobHandler) pause(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	jobID := req.URL.Query().Get("id")
	if jobID == "" {
		respond(res, "Job id required", false, http.StatusBadRequest)
		return
	}

	j, ok := JH.JobPool.Lookup(jobID)
	if !ok {
		respond(res, "Job id not found in running job pool or cache", false, http.StatusBadRequest)
		return
	}

	if j.State == job.Terminated || j.State == job.Done {
		respond(res, "Job has status "+string(j.State)+", cannot pause", false, http.StatusBadRequest)
		return
	}

	j.Pause()
	respond(res, "job paused successfully", true, http.StatusOK)
}

// resume - Handler for /resume
// Resumes a job given jobi
func (JH *JobHandler) resume(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	jobID := req.URL.Query().Get("id")
	if jobID == "" {
		respond(res, "Job id required", false, http.StatusBadRequest)
		return
	}

	j, ok := JH.JobPool.Lookup(jobID)
	if !ok {
		respond(res, "Job id not found in running job pool or cache", false, http.StatusBadRequest)
		return
	}
	if j.State == job.Terminated || j.State == job.Done {
		respond(res, "Job has status "+string(j.State)+", cannot resume", false, http.StatusBadRequest)
		return
	}

	j.Resume()
	respond(res, "job resumed successfully", true, http.StatusOK)
}

// terminate - Handler for /terminate
// terminate the task with given id
func (JH *JobHandler) terminate(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	jobID := req.URL.Query().Get("id")
	if jobID == "" {
		respond(res, "Job id required", false, http.StatusBadRequest)
		return
	}

	j, ok := JH.JobPool.Lookup(jobID)
	if !ok {
		respond(res, "Job id not found in running job pool or cache", false, http.StatusBadRequest)
		return
	}

	if j.State == job.Terminated || j.State == job.Done {
		respond(res, "Job has status "+string(j.State)+", cannot Terminate", false, http.StatusBadRequest)
		return
	}

	j.Terminate()
	respond(res, "job paused successfully", true, http.StatusOK)
}

// getStatus - Handler for /status
// Gives the status of a specific job
func (JH *JobHandler) getStatus(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	jobID := req.URL.Query().Get("id")
	if jobID == "" {
		respond(res, "Job id required", false, http.StatusBadRequest)
		return
	}

	j, ok := JH.JobPool.Lookup(jobID)
	if !ok {
		respond(res, "Job id not found in running job pool or cache", false, http.StatusBadRequest)
		return
	}
	respond(res, string(j.State), true, http.StatusOK)
}
