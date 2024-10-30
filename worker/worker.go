package worker

import (
	"sync"
	"time"
)

const (
	sleepTime = 500 * time.Millisecond
)

type Worker struct {
	id   int
	done chan struct{}
}

func New(id int) Worker {
	return Worker{
		id:   id,
		done: make(chan struct{}, 1),
	}
}

func (p Worker) Work(wg *sync.WaitGroup, input <-chan string, output chan<- string) {
	defer func() { wg.Done() }()
	for {
		select {
		case <-p.done:
			return
		default:
			select {
			case arg, ok := <-input:
				if !ok {
					return
				}
				time.Sleep(sleepTime)
				output <- Job(arg, p.id)
			default:
			}
		}
	}
}

func (p Worker) Deactivate() {
	p.done <- struct{}{}
}
