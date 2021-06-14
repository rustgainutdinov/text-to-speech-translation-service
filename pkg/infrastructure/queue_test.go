package infrastructure

import (
	"fmt"
	"testing"
	"text-to-speech-translation-service/pkg/domain"
	"time"
)

func TestQueue_AddTask(t *testing.T) {
	textToSpeechService := domain.NewTextToSpeechService(&mockTranslationTextToSpeechRepo{}, &mockExternalTextToSpeech{})
	queue := NewQueue(textToSpeechService)
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

type mockExternalTextToSpeech struct{}

func (t *mockExternalTextToSpeech) Translate(text string) (string, error) {
	return "olala", nil
}

type mockTranslationTextToSpeechRepo struct {
	translations []domain.Translation
}

func (m *mockTranslationTextToSpeechRepo) Store(translation domain.Translation) error {
	return nil
}

func (m *mockTranslationTextToSpeechRepo) FindOne(translationID domain.TranslationID) (domain.Translation, error) {
	return domain.Translation{}, nil
}
