package workerpool

import (
	"sync"
	"worker-pool/worker"
)

type (
	DataChan = chan string
)

type WorkerPool struct {
	workers []worker.Worker
	input   DataChan
	output  DataChan
	wg      sync.WaitGroup
	mu      sync.Mutex
}

func New(wcount int) WorkerPool {
	workers := make([]worker.Worker, wcount)
	for i := range workers {
		workers[i] = worker.New(i)
	}
	return WorkerPool{
		workers: workers,
		input:   make(DataChan),
		output:  make(DataChan),
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
	wp.mu.Lock()
	defer func() { wp.mu.Unlock() }()

	newWorkers := make([]worker.Worker, count)
	for i := range newWorkers {
		w := worker.New(i + len(wp.workers))

		wp.wg.Add(1)
		go w.Work(&wp.wg, wp.input, wp.output)

		newWorkers[i] = w
	}

	wp.workers = append(wp.workers, newWorkers...)
}

func (wp *WorkerPool) Remove(count int) {
	wp.mu.Lock()
	defer func() { wp.mu.Unlock() }()

	count = min(count, len(wp.workers))

	for i := len(wp.workers) - count; i < len(wp.workers); i++ {
		go wp.workers[i].Deactivate()
	}

	wp.workers = wp.workers[:len(wp.workers)-count]
}
