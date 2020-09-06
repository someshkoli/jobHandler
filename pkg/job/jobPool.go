package job

// Pool - pool of job used to store job with their uuids
type Pool map[string]*Job

// Lookup - check if element is present in the pool
func (P Pool) Lookup(id string) (*Job, bool) {
	j, ok := P[id]
	return j, ok
}

// Add - Adds the job in the pool
func (P Pool) Add(id string, job *Job) *Job {
	P[id] = job
	return job
}

// Remove - Removes a job from the pool
func (P Pool) Remove(id string) {
	delete(P, id)
}
