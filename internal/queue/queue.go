package queue

import (
	"github.com/rs/zerolog"
	"time"
)

type Data struct {
	TTL int
	F   func() bool
}

type Queue struct {
	logger     zerolog.Logger
	defaultTTL int
	data       []Data
}

func (q *Queue) Add(data Data) {
	q.data = append(q.data, data)
}

func (q *Queue) Run() {
	q.logger.Debug().Msg("starting queue worker")
	for {
		for _, data := range q.data {
			q.logger.Debug().Msg("item fetched and removed from queue")
			// Discard top element
			q.data = q.data[1:]

			q.logger.Debug().Msg("executing function")
			ok := data.F()
			if !ok {
				q.logger.Debug().Msg("function was not successfully executed")
				data.TTL -= 1
				if data.TTL <= 0 {
					q.logger.Warn().Msg("TTL exceed. Item is discarded")
					continue
				}

				q.logger.Debug().Msgf("Reschedule item, remaining TTL: %d", data.TTL)
				q.data = append(q.data, data)
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func New(defaultTTL int, logger zerolog.Logger) *Queue {
	var data []Data
	return &Queue{
		logger:     logger,
		defaultTTL: defaultTTL,
		data:       data,
	}
}
