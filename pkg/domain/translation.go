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

type TranslationRepo interface {
	Store(translation Translation) error
	FindOne(translationID TranslationID) (Translation, error)
}

type TranslationID uuid.UUID

type Translation struct {
	ID         TranslationID
	UserID     uuid.UUID
	Text       string
	Status     int
	SpeechData string
}
