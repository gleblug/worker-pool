package workerpool

type WorkerPool struct {
	workersCount int
	workers      []Worker
	input        chan string
	output       chan string
	running      bool
}

func New(wcount int) WorkerPool {
	workers := make([]Worker, wcount)
	for i := range workers {
		workers[i] = NewWorker(i)
	}
	return WorkerPool{
		workersCount: wcount,
		workers:      workers,
		input:        make(chan string, wcount),
		output:       make(chan string, wcount),
		running:      false,
	}
}

func (wp *WorkerPool) Run() {
	// TODO: add goroutine that print output
	defer func() { wp.running = true }()
	for _, w := range wp.workers {
		go w.Work(wp.input, wp.output)
	}
}

func (wp *WorkerPool) Stop() {
	defer func() { wp.running = false }()
	for _, w := range wp.workers {
		go w.Deactivate()
	}
}

func (wp WorkerPool) Count() int {
	return wp.workersCount
}

func (wp WorkerPool) Running() bool {
	return wp.running
}

func (wp *WorkerPool) Remove(count int) {
	// TODO: add count < workersCount check
	defer func() { wp.workersCount -= count }()
	if wp.Running() {
		for i := wp.workersCount - count; i < wp.workersCount; i++ {
			// TODO: wait until workers is really deactivate (wait group)
			go wp.workers[i].Deactivate()
		}
	}
	wp.workers = wp.workers[:wp.workersCount-count]
}

func (wp *WorkerPool) Add(count int) {
	defer func() { wp.workersCount += count }()
	newWorkers := make([]Worker, count)
	for i := range newWorkers {
		worker := NewWorker(i + wp.workersCount)
		newWorkers[i] = worker
	}
	if wp.Running() {
		for _, w := range newWorkers {
			go w.Work(wp.input, wp.output)
		}
	}
	wp.workers = append(wp.workers, newWorkers...)
}
