package job

import (
	"fmt"
	"time"
)

func (J *Job) updateStatus(status Status) {
	J.Lock()
	J.State = status
	J.Unlock()
}

// Start - Starts the job
func (J *Job) Start() {
	J.updateStatus(Running)
	go func() {
		fmt.Printf("[STARTED][%s][%s]\n", J.ID, time.Now().Format("2 Jan 2006 15:04:05"))

		for _, d := range J.Data {
			select {
			case <-J.terminate:
				J.Terminate()
				return
			case <-J.pause:
				select {
				case <-J.terminate:
					J.Terminate()
					return
				case <-J.resume:
				}
			default:
				if writeData(J, d) != nil {
					fmt.Printf("[DATA_STORE_MISS][%s][%s] - %s\n", J.ID, time.Now().Format("2 Jan 2006 15:04:05"), d)
					continue
				}
			}
		}
		J.done()
	}()
}

// Pause - pause the job and updates status
func (J *Job) Pause() {
	if J.State != Running {
		return
	}

	J.updateStatus(Paused)
	J.pause <- struct{}{}

	fmt.Printf("[PAUSED][%s][%s]\n", J.ID, time.Now().Format("2 Jan 2006 15:04:05"))
	return
}

// Resume Job and updates status
func (J *Job) Resume() {
	if J.State != Paused {
		return
	}

	J.updateStatus(Running)
	J.resume <- struct{}{}

	fmt.Printf("[RESUMED][%s][%s]\n", J.ID, time.Now().Format("2 Jan 2006 15:04:05"))
}

// Terminate - Terminates the Job and does the cleanup operation
func (J *Job) Terminate() {
	if J.State != Running && J.State != Paused {
		return
	}

	J.updateStatus(Terminated)
	clean(J)

	fmt.Printf("[TERMIATED][%s][%s]\n", J.ID, time.Now().Format("2 Jan 2006 15:04:05"))
}

func (J *Job) done() {
	J.updateStatus(Done)
	clean(J)
	fmt.Printf("[Done][%s][%s]\n", J.ID, time.Now().Format("2 Jan 2006 15:04:05"))
}
