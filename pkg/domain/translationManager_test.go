package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

const mockTranslationSpeechToText = "ya perevel"

func TestTranslationService_AddTextToSpeechTranslation(t *testing.T) {
	repo := mockTranslationTextToSpeechRepo{}
	translationService := NewTranslationManager(&mockTranslationQueue{}, &repo, &mockBalanceService{})
	textToTranslate := "Hello, world"
	translationID, err := translationService.AddTranslation(textToTranslate, uuid.New())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.translations))
	translation, err := repo.FindOne(translationID)
	assert.Nil(t, err)
	assert.Equal(t, textToTranslate, translation.Text)
}

func TestTranslationService_TranslateTextToSpeech(t *testing.T) {
	repo := mockTranslationTextToSpeechRepo{}
	translationService := NewTranslationManager(&mockTranslationQueue{}, &repo, &mockBalanceService{})
	textToTranslate := "Hello, world"
	translationID, err := translationService.AddTranslation(textToTranslate, uuid.New())
	assert.Nil(t, err)
	textToSpeechService := NewTextToSpeechService(&repo, &mockExternalTextToSpeech{}, &mockExternalEventBroker{})
	err = textToSpeechService.Translate(translationID)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.translations))
	translation, err := repo.FindOne(translationID)
	assert.Nil(t, err)
	assert.Equal(t, textToTranslate, translation.Text)
	assert.Equal(t, mockTranslationSpeechToText, translation.SpeechData)
}

type mockTranslationTextToSpeechRepo struct {
	translations []Translation
}

func (m *mockTranslationTextToSpeechRepo) Store(translation Translation) error {
	for i, repoTranslation := range m.translations {
		if repoTranslation.ID == translation.ID {
			m.translations[i] = translation
			return nil
		}
	}
	m.translations = append(m.translations, translation)
	return nil
}

func (m *mockTranslationTextToSpeechRepo) FindOne(translationID TranslationID) (Translation, error) {
	for _, translation := range m.translations {
		if translation.ID == translationID {
			return translation, nil
		}
	}
	return Translation{}, ErrTranslationIsNotFound
}

type mockExternalTextToSpeech struct{}

func (t *mockExternalTextToSpeech) Translate(text string) (string, error) {
	return mockTranslationSpeechToText, nil
}

type mockTranslationQueue struct{}

func (t *mockTranslationQueue) AddTask(task Task) {}

func (t *mockTranslationQueue) Start() {}

type mockBalanceService struct{}

func (t *mockBalanceService) CanWriteOf(userID uuid.UUID, amountOfSymbols int) (bool, error) {
	return true, nil
}

type mockExternalEventBroker struct{}

func (t *mockExternalEventBroker) TextTranslated(userID uuid.UUID, amountOfSymbols int) error {
	return nil
}
