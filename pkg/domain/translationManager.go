package domain

import (
	"github.com/google/uuid"
)

type TranslationManager interface {
	AddTranslation(text string, userID uuid.UUID) (TranslationID, error)
	SaveTranslatedData(translationID TranslationID, translatedData string) error
}

type translationManager struct {
	translationRepo TranslationRepo
}

func (t *translationManager) AddTranslation(text string, userID uuid.UUID) (TranslationID, error) {
	translationID := TranslationID(uuid.New())
	translation := Translation{
		ID:     translationID,
		UserID: userID,
		Text:   text,
		Status: TranslationStatusWaiting,
	}
	err := t.translationRepo.Store(translation)
	if err != nil {
		return TranslationID{}, err
	}
	return translationID, nil
}

func (t *translationManager) SaveTranslatedData(translationID TranslationID, translatedData string) error {
	translation, err := t.translationRepo.FindOne(translationID)
	if err != nil {
		return err
	}
	return t.translationRepo.Store(Translation{
		ID:         translationID,
		UserID:     translation.UserID,
		Text:       translation.Text,
		Status:     TranslationStatusSuccess,
		SpeechData: translation.SpeechData,
	})
}

func NewTranslationManager(translationRepo TranslationRepo) TranslationManager {
	return &translationManager{
		translationRepo: translationRepo,
	}
}
