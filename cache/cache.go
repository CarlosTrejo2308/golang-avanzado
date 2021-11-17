package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Printf("Calculating fibonacci(%d)\n", n)
	time.Sleep(time.Second * 5)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	// Lock the service
	s.Lock.RLock()

	// If exist, then wait for the respons of the existing worker
	if s.InProgress[job] {
		s.Lock.RUnlock()

		// This is where we going to get the response
		response := make(chan int)
		defer close(response)

		// Add the channel to the pending list
		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Waiting for job %d\n", job)

		// Wait for the response
		res := <-response
		fmt.Printf("Got response %d\n", res)
		return
	}

	// If not exist, then start the job
	s.Lock.RUnlock()

	// We are working on it
	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Working on job %d\n", job)
	res := ExpensiveFibonacci(job)

	// Get the pending workers for this job
	s.Lock.RLock()
	pendingWorkers, exist := s.IsPending[job]
	s.Lock.RUnlock()

	// If there are pending workers, then send the response
	if exist {
		for _, worker := range pendingWorkers {
			worker <- res
		}
	}

	// We are done with this job, reset the state
	s.Lock.Lock()
	delete(s.InProgress, job)
	delete(s.IsPending, job)
	//s.InProgress[job] = false
	//s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()

	fmt.Printf("Finished job %d, got %d\n", job, res)
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{1, 2, 3, 4, 4, 5, 5, 5, 8, 8, 8, 8, 8}

	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, job := range jobs {

		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(job)
	}
	wg.Wait()

}
