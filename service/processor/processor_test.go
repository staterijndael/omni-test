package processor

import (
	"context"
	"testing"
	"time"

	"omni-test/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetLimits() (uint64, time.Duration) {
	args := m.Called()
	return args.Get(0).(uint64), args.Get(1).(time.Duration)
}

func (m *MockService) Process(ctx context.Context, batch service.Batch) error {
	args := m.Called(ctx, batch)
	return args.Error(0)
}

func newMockService(limit uint64, duration time.Duration, err error) *MockService {
	m := &MockService{}
	m.On("GetLimits").Return(limit, duration)
	m.On("Process", mock.Anything, mock.Anything).Return(err)
	return m
}

func TestProcessBatch(t *testing.T) {
	tests := []struct {
		name               string
		batchSize          int
		limit              uint64
		duration           time.Duration
		processError       error
		expectedError      error
		expectedBlockCount uint64
	}{
		{
			name:               "Normal Processing",
			batchSize:          3,
			limit:              5,
			duration:           time.Second,
			processError:       nil,
			expectedError:      nil,
			expectedBlockCount: 3,
		},
		{
			name:               "Blocked Processing",
			batchSize:          5,
			limit:              5,
			duration:           time.Second,
			processError:       service.ErrBlocked,
			expectedError:      service.ErrBlocked,
			expectedBlockCount: 0,
		},
		{
			name:               "Processing Limit",
			batchSize:          15,
			limit:              5,
			duration:           time.Second,
			processError:       nil,
			expectedError:      service.ErrBlocked,
			expectedBlockCount: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockService := newMockService(test.limit, test.duration, test.processError)
			sut := NewProcessor(mockService)
			batch := make(service.Batch, test.batchSize)

			err := sut.ProcessBatch(context.Background(), batch)

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedBlockCount, sut.blockingProcessedItemsCount)
		})
	}
}
