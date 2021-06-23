package service

import (
	"fmt"
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/app"
	"text-to-speech-translation-service/pkg/app/dataProvider"
	"text-to-speech-translation-service/pkg/domain"
)

type TranslationService interface {
	AddTextToTranslate(userID uuid.UUID, text string) (uuid.UUID, error)
	GetTranslationData(translationID uuid.UUID) (string, error)
	GetTranslationStatus(translationID uuid.UUID) (int, error)
}

type BalanceService interface {
	CanWriteOf(userID uuid.UUID, score int) (bool, error)
}

type translationService struct {
	unitOfWorkFactory       dataProvider.UnitOfWorkFactory
	translationQueue        app.Queue
	balanceService          BalanceService
	translationQueryService dataProvider.TranslationQueryService
}

var ErrThereAreNotEnoughSymbolsToWriteOff = fmt.Errorf("there are not enough symbols to write off")

func (b *translationService) AddTextToTranslate(userID uuid.UUID, text string) (uuid.UUID, error) {
	canWriteOf, err := b.balanceService.CanWriteOf(userID, len(text))
	if err != nil {
		return uuid.UUID{}, err
	}
	if !canWriteOf {
		return uuid.UUID{}, ErrThereAreNotEnoughSymbolsToWriteOff
	}
	var translationID domain.TranslationID
	err = b.unitOfWorkFactory.NewUnitOfWork(func(provider dataProvider.RepositoryProvider) error {
		translationID, err = domain.NewTranslationManager(provider.TranslationRepo()).AddTextToTranslate(text, userID)
		if err != nil {
			return err
		}
		return nil
	})
	b.translationQueue.AddTask(app.Task{
		Type: app.TextTranslatedTaskType,
		Data: app.TextTranslated{
			TranslationID: uuid.UUID(translationID),
			Text:          text,
		},
	})
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

func NewTranslationService(unitOfWorkFactory dataProvider.UnitOfWorkFactory, translationQueue app.Queue, balanceService BalanceService, translationQueryService dataProvider.TranslationQueryService) TranslationService {
	return &translationService{
		unitOfWorkFactory:       unitOfWorkFactory,
		translationQueue:        translationQueue,
		balanceService:          balanceService,
		translationQueryService: translationQueryService,
	}
}
