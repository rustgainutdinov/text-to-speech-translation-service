package infrastructure

import (
	"fmt"
	"testing"
	"text-to-speech-translation-service/pkg/domain"
	"time"
)

func TestQueue_AddTask(t *testing.T) {
	queue := NewQueue()
	go queue.Start()
	go func() {
		time.Sleep(time.Second * 2)
		queue.AddTask(domain.Task{
			TranslationID: domain.TranslationID{},
			Text:          "msg 1",
		})
	}()
	queue.AddTask(domain.Task{
		TranslationID: domain.TranslationID{},
		Text:          "msg 2",
	})
	queue.AddTask(domain.Task{
		TranslationID: domain.TranslationID{},
		Text:          "msg 3",
	})
	queue.AddTask(domain.Task{
		TranslationID: domain.TranslationID{},
		Text:          "msg 4",
	})
	fmt.Println("end of test")
}
