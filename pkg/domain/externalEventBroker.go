package domain

import "github.com/google/uuid"

type ExternalEventBroker interface {
	TextTranslated(userID uuid.UUID, amountOfSymbols int) error
}
