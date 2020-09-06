package job

import (
	"fmt"
	"time"
)

func clean(j *Job) {
	j.Lock()
	close(j.pause)
	close(j.resume)
	close(j.terminate)
	j.Unlock()
}

// TODO: Check for persistent in future
func writeData(J *Job, data string) error {
	fmt.Printf("[WRITING][%s][%s] - %s\n", J.ID, time.Now().Format("2 Jan 2006 15:04:05"), data)
	// TODO: update data to db here and fetch in resume method
	time.Sleep(time.Duration(1000) * time.Millisecond)
	return nil
}
