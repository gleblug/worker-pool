package workerpool

type WorkerPool struct {
	input chan string
	output chan string
	workersCount int

}

func New() WorkerPool {

}

func (wp WorkerPool) Run() {

}

func (wp WorkerPool) Count() int {
	return wp.workersCount
}

func (wp WorkerPool) Remove(count int) {

}

func (wp WorkerPool) Add(count int) {

}