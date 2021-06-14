package app

import (
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/domain"
)

type TranslationService interface {
	Translate(userID uuid.UUID, text string) (uuid.UUID, error)
}

type translationService struct {
	translationRepo  domain.TranslationRepo
	translationQueue domain.TranslationQueue
	balanceService   domain.BalanceService
}

func (b *translationService) Translate(userID uuid.UUID, text string) (uuid.UUID, error) {
	translationID, err := domain.NewTranslationManager(b.translationQueue, b.translationRepo, b.balanceService).AddTranslation(text, userID)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.UUID(translationID), nil
}

func NewTranslationService(translationRepo domain.TranslationRepo, translationQueue domain.TranslationQueue, balanceService domain.BalanceService) TranslationService {
	return &translationService{
		translationRepo:  translationRepo,
		translationQueue: translationQueue,
		balanceService:   balanceService,
	}
}
