package app

import (
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/domain"
)

type TranslationService interface {
	Translate(userID uuid.UUID, text string) (uuid.UUID, error)
	GetTranslationData(translationID uuid.UUID) (string, error)
	GetTranslationStatus(translationID uuid.UUID) (int, error)
}

type translationService struct {
	unitOfWorkFactory       domain.UnitOfWorkFactory
	translationQueue        domain.TranslationQueue
	balanceService          domain.BalanceService
	translationQueryService TranslationQueryService
}

func (b *translationService) Translate(userID uuid.UUID, text string) (uuid.UUID, error) {
	translationID, err := domain.NewTranslationManager(b.translationQueue, b.unitOfWorkFactory, b.balanceService).AddTranslation(text, userID)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.UUID(translationID), nil
}

func (b *translationService) GetTranslationData(translationID uuid.UUID) (string, error) {
	translationDTO, err := b.translationQueryService.GetTranslationData(translationID)
	if err != nil {
		return "", err
	}
	return translationDTO.TranslatedData(), nil
}

func (b *translationService) GetTranslationStatus(translationID uuid.UUID) (int, error) {
	translationDTO, err := b.translationQueryService.GetTranslationData(translationID)
	if err != nil {
		return 0, err
	}
	return translationDTO.Status(), nil
}

func NewTranslationService(unitOfWorkFactory domain.UnitOfWorkFactory, translationQueue domain.TranslationQueue, balanceService domain.BalanceService, translationQueryService TranslationQueryService) TranslationService {
	return &translationService{
		unitOfWorkFactory:       unitOfWorkFactory,
		translationQueue:        translationQueue,
		balanceService:          balanceService,
		translationQueryService: translationQueryService,
	}
}
