package dataProvider

import (
	"fmt"
	"github.com/google/uuid"
)

var ErrTranslationIsNotFound = fmt.Errorf("translation is not found")

type TranslationQueryService interface {
	GetTranslationData(translationID uuid.UUID) (TranslationDTO, error)
}

type TranslationDTO interface {
	Status() int
	TranslatedData() string
	UserID() string
	Text() string
}
