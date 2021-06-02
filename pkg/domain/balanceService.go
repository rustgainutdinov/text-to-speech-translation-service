package domain

import "github.com/google/uuid"

type BalanceService interface {
	CanWriteOf(userID uuid.UUID, amountOfSymbols int) (bool, error)
}
