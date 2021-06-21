package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranslationService_AddTextToSpeechTranslation(t *testing.T) {
	repo := mockTranslationTextToSpeechRepo{}
	translationService := NewTranslationManager(&repo)
	textToTranslate := "Hello, world"
	translationID, err := translationService.AddTranslation(textToTranslate, uuid.New())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.translations))
	translation, err := repo.FindOne(translationID)
	assert.Nil(t, err)
	assert.Equal(t, textToTranslate, translation.Text)
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
