package workerpool

import (
	"fmt"
	"sync"
	"testing"
	"time"
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

func TestWorkerPoolAdding(t *testing.T) {
	const initCount = 1

	wp := New(initCount)
	t.Logf("Worker pool with %d worker created", initCount)
	go wp.Run()
	t.Log("Worker pool running")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go generateInput(&wg, dataCount, wp.input)
	go printOutput(&wg, wp.output)

	time.Sleep(2 * time.Second)
	wp.Add(1)
	t.Log("1 worker added")

	time.Sleep(2 * time.Second)
	wp.Add(2)
	t.Log("2 workers added")

	wg.Wait()
}

func TestWorkerPoolRemoving(t *testing.T) {
	const initCount = 4

	wp := New(initCount)
	t.Logf("Worker pool with %d worker created", initCount)
	go wp.Run()
	t.Log("Worker pool running")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go generateInput(&wg, dataCount, wp.input)
	go printOutput(&wg, wp.output)

	time.Sleep(2 * time.Second)
	wp.Remove(2)
	t.Log("2 workers removed")

	time.Sleep(2 * time.Second)
	wp.Remove(1)
	t.Log("1 workers removed")

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
