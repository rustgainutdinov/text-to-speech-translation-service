package queue

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/app"
	"time"
)

type queue struct {
	in                  chan app.Task
	ctx                 context.Context
	textToSpeechService app.TextToSpeechService
}

func (s *queue) AddTask(task app.Task) {
	s.in <- task
}

func (s *queue) Start() {
	fmt.Println("channel started")
	for {
		select {
		case task := <-s.in:
			{
				err := s.textToSpeechService.Translate(uuid.UUID(task.TranslationID))
				if err != nil {
					fmt.Println(err)
				}
				time.Sleep(1 * time.Second)
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func NewQueue(textToSpeechService app.TextToSpeechService) app.TranslationQueue {
	q := queue{
		in:                  make(chan app.Task),
		ctx:                 context.Background(),
		textToSpeechService: textToSpeechService,
	}
	go q.Start()
	return &q
}
