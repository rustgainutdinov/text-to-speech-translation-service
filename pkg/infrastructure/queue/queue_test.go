package queue

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
	"text-to-speech-translation-service/pkg/app"
	"text-to-speech-translation-service/pkg/domain"
	"time"
)

func TestQueue_AddTask(t *testing.T) {
	textToSpeechService := app.NewTextToSpeechService(&mockTranslationTextToSpeechRepo{}, &mockExternalTextToSpeech{}, &mockExternalEventBroker{})
	queue := NewQueue(textToSpeechService)
	go func() {
		time.Sleep(time.Second * 2)
		queue.AddTask(app.Task{
			TranslationID: domain.TranslationID{},
			Text:          "msg 1",
		})
	}()
	queue.AddTask(app.Task{
		TranslationID: domain.TranslationID{},
		Text:          "msg 2",
	})
	queue.AddTask(app.Task{
		TranslationID: domain.TranslationID{},
		Text:          "msg 3",
	})
	queue.AddTask(app.Task{
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

type mockExternalEventBroker struct{}

func (t *mockExternalEventBroker) TextTranslated(userID uuid.UUID, score int) error {
	return nil
}
