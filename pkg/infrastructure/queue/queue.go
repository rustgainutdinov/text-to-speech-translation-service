package queue

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"text-to-speech-translation-service/pkg/app"
	"text-to-speech-translation-service/pkg/app/service"
	"time"
)

type queue struct {
	in                  chan app.Task
	ctx                 context.Context
	textToSpeechService service.TextToSpeechService
}

func (s *queue) AddTask(task app.Task) {
	go func() {
		s.in <- task
	}()
	log.Info("task added successfully")
}

var ErrQueueTaskDataParsing = fmt.Errorf("can't parse queue task data")
var ErrUnknownTaskType = fmt.Errorf("unknown task type")

func (s *queue) Start() {
	for {
		select {
		case task := <-s.in:
			{
				var err error
				switch task.Type {
				case app.TextTranslatedTaskType:
					err = s.textTranslatedTaskHandler(task)
				default:
					err = ErrUnknownTaskType
				}
				if err != nil {
					log.WithFields(log.Fields{"task": task}).Error(err)
				}
				log.Info("task handled successfully")
				time.Sleep(1 * time.Second)
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func (s *queue) textTranslatedTaskHandler(task app.Task) error {
	value, ok := task.Data.(app.TextTranslated)
	if !ok {
		return ErrQueueTaskDataParsing
	}
	return s.textToSpeechService.Translate(value.TranslationID)
}

func NewQueue(textToSpeechService service.TextToSpeechService) app.Queue {
	q := queue{
		in:                  make(chan app.Task),
		ctx:                 context.Background(),
		textToSpeechService: textToSpeechService,
	}
	go q.Start()
	return &q
}
