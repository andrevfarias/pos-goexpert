package usecase

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type PerformStressTestUsecaseInput struct {
	URL         string
	Requests    int
	Concurrency int
}

type PerformStressTestUsecaseOutput struct {
	Requests       int
	FailedRequests int
	StatusCode     []int
	TestDuration   time.Duration
}

type PerformStressTestUsecaseInterface interface {
	Execute(input PerformStressTestUsecaseInput) (PerformStressTestUsecaseOutput, error)
}

type PerformStressTestUsecase struct{}

func NewPerformStressTestUsecase() PerformStressTestUsecaseInterface {
	return &PerformStressTestUsecase{}
}

func (u *PerformStressTestUsecase) Execute(input PerformStressTestUsecaseInput) (PerformStressTestUsecaseOutput, error) {
	startTime := time.Now()

	output := PerformStressTestUsecaseOutput{}

	workers := make(chan struct{}, input.Concurrency)

	defer close(workers)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < input.Requests; i++ {
		workers <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() { <-workers }()
			defer wg.Done()
			res, err := http.Get(input.URL)

			mu.Lock()
			output.Requests++
			if err != nil {
				output.FailedRequests++
				fmt.Printf("Request error -> %v\n", err)
			} else {
				defer res.Body.Close()
				output.StatusCode = append(output.StatusCode, res.StatusCode)
			}
			mu.Unlock()
		}()
	}

	wg.Wait()

	output.TestDuration = time.Since(startTime)
	return output, nil
}
