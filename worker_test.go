package workerpool

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

const (
	testCount = 3
	maxIndex  = 1000
)

func TestWorker(t *testing.T) {
	input := make(chan string)
	output := make(chan string)
	wg := sync.WaitGroup{}
	for range testCount {
		workerIndex := rand.Intn(maxIndex)
		w := NewWorker(workerIndex)
		t.Log("Worker created")

		wg.Add(1)
		go w.Work(&wg, input, output)
		t.Log("Worker's goroutine started")
		input <- mockString

		expect := fmt.Sprintf("%d work: %s", workerIndex, mockString)
		if res := <-output; expect != res {
			t.Errorf("Expected %q != Result %q", expect, res)
		}

		t.Log("Expected result passed")
		w.Deactivate()
		t.Log("Worker deactivated")
	}
}
