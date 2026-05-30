package scheduler

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type JobFunc func(context.Context) error

type Scheduler struct {
	scheduler gocron.Scheduler
	ctx       context.Context
}

func New(ctx context.Context) (*Scheduler, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &Scheduler{
		scheduler: s,
		ctx:       ctx,
	}, nil
}

func (s *Scheduler) Every(duration time.Duration, job JobFunc) error {
	_, err := s.scheduler.NewJob(
		gocron.DurationJob(duration),
		gocron.NewTask(func() {
			_ = job(s.ctx)
		}),
	)

	return err
}

func (s *Scheduler) Start() {
	s.scheduler.Start()
}

func (s *Scheduler) Stop() error {
	return s.scheduler.Shutdown()
}
