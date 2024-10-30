package worker

import "fmt"

func Job(arg string, workerId int) string {
	return fmt.Sprintf("worker %d works: %q", workerId, arg)
}