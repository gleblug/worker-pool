package workerpool

import (
	"fmt"
	"sync"
	"time"
)

const (
	sleepTime = 1 * time.Second
)

type Worker struct {
	id   int
	done chan struct{}
}

func NewWorker(id int) Worker {
	return Worker{
		id:   id,
		done: make(chan struct{}),
	}
}

func (p Worker) Work(wg *sync.WaitGroup, input <-chan string, output chan<- string) {
	defer func() { wg.Done() }()
	for {
		select {
		case arg, ok := <-input:
			if !ok {
				return
			}
			output <- fmt.Sprintf("%d work: %s", p.id, arg)
			time.Sleep(sleepTime)
		case <-p.done:
			return
		}
	}
}

func (p *Worker) Deactivate() {
	p.done <- struct{}{}
}
