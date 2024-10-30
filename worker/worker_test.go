package worker

import (
	"math/rand"
	"sync"
	"testing"
)

const (
	workersCount = 3
	maxIndex     = 1000
	mockString   = "test_string"
)

func TestWorker(t *testing.T) {
	input := make(chan string)
	output := make(chan string)
	wg := sync.WaitGroup{}
	for range workersCount {
		workerIndex := rand.Intn(maxIndex)
		w := New(workerIndex)
		t.Log("Worker created")

		wg.Add(1)
		go w.Work(&wg, input, output)
		t.Log("Worker's goroutine started")
		input <- mockString

		expect := Job(mockString, workerIndex)
		if res := <-output; expect != res {
			t.Errorf("Expected %q != Result %q", expect, res)
		}
		t.Log("Expected result passed")

		w.Deactivate()
		t.Log("Worker deactivated")
	}
}
