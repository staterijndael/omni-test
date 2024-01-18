package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MockService struct {
	mu       sync.Mutex
	limit    uint64
	duration time.Duration
	counter  uint64
	lastTime time.Time
}

func NewMockService() *MockService {
	maxLimit := 100
	minLimit := 5

	limit := uint64(rand.Intn(maxLimit-minLimit) + minLimit)

	maxDuration := 20
	minDuration := 3

	duration := time.Duration(rand.Intn(maxDuration-minDuration+1)+minDuration) * time.Second

	fmt.Printf("Mock service duration: %.2fs, limit: %d\n", duration.Seconds(), limit)

	return &MockService{
		limit:    limit,
		duration: duration,
	}
}

func (s *MockService) GetLimits() (uint64, time.Duration) {
	return s.limit, s.duration
}

// Process processes the batch of items.
func (s *MockService) Process(_ context.Context, batch Batch) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if time.Since(s.lastTime) > s.duration {
		s.counter = 0
		s.lastTime = time.Now()
	}

	if s.counter >= s.limit {
		return errors.New("limit reached")
	}

	for _, item := range batch {
		_ = processItem(item)
		s.counter++
	}

	return nil
}

// processItem is a placeholder function for processing individual items.
func processItem(_ Item) error {
	return nil
}
