package infrastructure

import (
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/domain"
)

type balanceService struct {
	balanceServiceAddress string
}

func (b *balanceService) CanWriteOf(userID uuid.UUID, amountOfSymbols int) (bool, error) {
	return true, nil
}

func NewBalanceService(balanceServiceAddress string) domain.BalanceService {
	return &balanceService{balanceServiceAddress: balanceServiceAddress}
}
