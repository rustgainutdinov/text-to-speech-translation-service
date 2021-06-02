package infrastructure

import (
	"context"
	"fmt"
	"text-to-speech-translation-service/pkg/domain"
	"time"
)

type queue struct {
	in  chan domain.Task
	ctx context.Context
}

func (s *queue) AddTask(task domain.Task) {
	s.in <- task
}

func (s *queue) Start() {
	for {
		select {
		case task := <-s.in:
			{
				fmt.Println(task.Text)
				time.Sleep(3 * time.Second)
			}
		default:
			fmt.Println("default")
			time.Sleep(1 * time.Second)
		}
	}
}

func NewQueue() domain.TranslationQueue {
	sub := queue{
		in:  make(chan domain.Task),
		ctx: context.Background(),
	}
	return &sub
}
