package domain

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	TranslationStatusWaiting = iota
	TranslationStatusSuccess
	TranslationStatusError
)

var ErrTranslationIsNotFound = fmt.Errorf("translation is not found")

type TranslationTextToSpeechRepo interface {
	Store(translation TranslationTextToSpeech) error
	FindOne(translationID TranslationTextToSpeechID) (TranslationTextToSpeech, error)
}

type TranslationTextToSpeechID uuid.UUID

type TranslationTextToSpeech struct {
	ID         TranslationTextToSpeechID
	Text       string
	Status     int
	SpeechData string
}
