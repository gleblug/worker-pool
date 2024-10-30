package workerpool

import "fmt"

type Worker struct {
	id int
	done chan struct{}
}

func NewWorker(id int) Worker {
	return Worker{
		id: id,
		done: make(chan struct{}),
	}
}

func (p Worker) Work(input <-chan string, output chan<- string) {
	for {
		select {
		case arg, ok := <- input:
			if !ok {
				return
			}
			output <- fmt.Sprintf("%d work: %s", p.id, arg)
		case <-p.done:
			return
		}
	}
}

func (p *Worker) Deactivate() {
	p.done<-struct{}{}
}