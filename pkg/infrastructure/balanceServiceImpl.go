package infrastructure

import (
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/domain"
)

type balanceService struct{}

func (b *balanceService) CanWriteOf(userID uuid.UUID, amountOfSymbols int) (bool, error) {
	return true, nil
}

func NewBalanceService() domain.BalanceService {
	return &balanceService{}
}
