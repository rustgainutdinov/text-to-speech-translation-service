package app

import "github.com/google/uuid"

type BalanceService interface {
	CanWriteOf(userID uuid.UUID, score int) (bool, error)
}
