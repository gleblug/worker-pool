package workerpool

import "sync"

type WorkerPool struct {
	workers []Worker
	input   chan string
	output  chan string
	wg      sync.WaitGroup
	mu      sync.Mutex
}

func New(wcount int) WorkerPool {
	workers := make([]Worker, wcount)
	for i := range workers {
		workers[i] = NewWorker(i)
	}
	return WorkerPool{
		workers: workers,
		input:   make(chan string),
		output:  make(chan string),
		wg:      sync.WaitGroup{},
		mu:      sync.Mutex{},
	}
}

func (wp *WorkerPool) Run() {
	defer func() { close(wp.output) }()
	wp.mu.Lock()
	for _, w := range wp.workers {
		wp.wg.Add(1)
		go w.Work(&wp.wg, wp.input, wp.output)
	}
	wp.mu.Unlock()
	wp.wg.Wait()
}

func (wp *WorkerPool) Add(count int) {
	newWorkers := make([]Worker, count)
	for i := range newWorkers {
		w := NewWorker(i + len(wp.workers))

		wp.wg.Add(1)
		go w.Work(&wp.wg, wp.input, wp.output)

		newWorkers[i] = w
	}
	wp.mu.Lock()
	wp.workers = append(wp.workers, newWorkers...)
	wp.mu.Unlock()
}

// func (wp *WorkerPool) Remove(count int) {
// 	// TODO: add count < workers count check
// 	for i := wp.workersCount - count; i < wp.workersCount; i++ {
// 		// TODO: wait until workers is really deactivate (wait group)
// 		go wp.workers[i].Deactivate()
// 	}
// 	wp.workers = wp.workers[:wp.workersCount-count]
// }
