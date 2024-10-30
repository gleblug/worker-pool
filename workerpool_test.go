package workerpool

import (
	"fmt"
	"sync"
	"testing"
)

const (
	dataCount  = 25
	mockString = "test_string"
)

func TestWorkerPool(t *testing.T) {
	wp := New(10)
	t.Log("Worker pool created")
	go wp.Run()
	t.Log("Worker pool running")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go generateInput(&wg, dataCount, wp.input)
	go printOutput(&wg, wp.output)

	wg.Wait()
}

func generateInput(wg *sync.WaitGroup, count int, input chan<- string) {
	defer func() {
		close(input)
		wg.Done()
	}()
	for range count {
		input <- mockString
	}
}

func printOutput(wg *sync.WaitGroup, output <-chan string) {
	defer func() { wg.Done() }()
	for v := range output {
		fmt.Println(v)
	}
}
