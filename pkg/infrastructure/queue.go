package infrastructure

import (
	"context"
	"fmt"
	"text-to-speech-translation-service/pkg/domain"
	"time"
)

type queue struct {
	in                  chan domain.Task
	ctx                 context.Context
	textToSpeechService domain.TextToSpeechService
}

func (s *queue) AddTask(task domain.Task) {
	fmt.Println("task added")
	s.in <- task
}

func (s *queue) Start() {
	fmt.Println("queue started")
	for {
		select {
		case task := <-s.in:
			{
				err := s.textToSpeechService.Translate(task.TranslationID)
				if err != nil {
					fmt.Println(err)
				}
				time.Sleep(10 * time.Second)
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func NewQueue(textToSpeechService domain.TextToSpeechService) domain.TranslationQueue {
	q := queue{
		in:                  make(chan domain.Task),
		ctx:                 context.Background(),
		textToSpeechService: textToSpeechService,
	}
	go q.Start()
	return &q
}
