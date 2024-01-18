package processor

import (
	"context"
	"omni-test/service"
	"sync"
	"time"
)

type Processor struct {
	limit         uint64
	limitDuration time.Duration

	lastBlockingTime            time.Time
	blockingProcessedItemsCount uint64
	blockingMx                  sync.Mutex

	service service.Service
}

func NewProcessor(serv service.Service) *Processor {
	limit, limitDuration := serv.GetLimits()

	return &Processor{
		limit:         limit,
		limitDuration: limitDuration,

		blockingMx: sync.Mutex{},

		lastBlockingTime: time.Now(),

		service: serv,
	}
}

func (s *Processor) ProcessBatch(ctx context.Context, batch service.Batch) error {
	s.blockingMx.Lock()
	defer s.blockingMx.Unlock()

	// in this case, we do not take into account the network call error
	// as we would do in production code, since there is no external service
	if s.lastBlockingTime.Add(s.limitDuration).Before(time.Now()) {
		s.lastBlockingTime = time.Now()
		s.blockingProcessedItemsCount = 0
	}

	if s.blockingProcessedItemsCount+uint64(len(batch)) > s.limit {
		return service.ErrBlocked
	}

	err := s.service.Process(ctx, batch)
	if err != nil {
		s.lastBlockingTime = time.Now().Add(s.limitDuration)
		return err
	}

	s.blockingProcessedItemsCount += uint64(len(batch))

	return nil
}
