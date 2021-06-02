package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const mockTranslationSpeechToText = "ya perevel"

func TestTranslationService_AddTextToSpeechTranslation(t *testing.T) {
	repo := mockTranslationTextToSpeechRepo{}
	translationService := NewTranslationService(&mockTranslationQueue{}, &repo, &mockTextToSpeechService{})
	textToTranslate := "Hello, world"
	translationID, err := translationService.AddTextToSpeechTranslation(textToTranslate)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.translations))
	translation, err := repo.FindOne(TranslationTextToSpeechID(translationID))
	assert.Nil(t, err)
	assert.Equal(t, textToTranslate, translation.Text)
}

func TestTranslationService_TranslateTextToSpeech(t *testing.T) {
	repo := mockTranslationTextToSpeechRepo{}
	translationService := NewTranslationService(&mockTranslationQueue{}, &repo, &mockTextToSpeechService{})
	textToTranslate := "Hello, world"
	translationID, err := translationService.AddTextToSpeechTranslation(textToTranslate)
	assert.Nil(t, err)
	err = translationService.TranslateTextToSpeech(TranslationTextToSpeechID(translationID))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.translations))
	translation, err := repo.FindOne(TranslationTextToSpeechID(translationID))
	assert.Nil(t, err)
	assert.Equal(t, textToTranslate, translation.Text)
	assert.Equal(t, mockTranslationSpeechToText, translation.SpeechData)
}

type mockTranslationTextToSpeechRepo struct {
	translations []TranslationTextToSpeech
}

func (m *mockTranslationTextToSpeechRepo) Store(translation TranslationTextToSpeech) error {
	for i, repoTranslation := range m.translations {
		if repoTranslation.ID == translation.ID {
			m.translations[i] = translation
			return nil
		}
	}
	m.translations = append(m.translations, translation)
	return nil
}

func (m *mockTranslationTextToSpeechRepo) FindOne(translationID TranslationTextToSpeechID) (TranslationTextToSpeech, error) {
	for _, translation := range m.translations {
		if translation.ID == translationID {
			return translation, nil
		}
	}
	return TranslationTextToSpeech{}, ErrTranslationIsNotFound
}

type mockTextToSpeechService struct{}

func (t *mockTextToSpeechService) Translate(text string) (string, error) {
	return mockTranslationSpeechToText, nil
}

type mockTranslationQueue struct{}

func (t *mockTranslationQueue) AddJob(task Task) {
}
