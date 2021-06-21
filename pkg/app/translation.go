package app

import (
	"fmt"
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/domain"
)

type TranslationService interface {
	Translate(userID uuid.UUID, text string) (uuid.UUID, error)
	GetTranslationData(translationID uuid.UUID) (string, error)
	GetTranslationStatus(translationID uuid.UUID) (int, error)
}

type translationService struct {
	unitOfWorkFactory       UnitOfWorkFactory
	translationQueue        TranslationQueue
	balanceService          BalanceService
	translationQueryService TranslationQueryService
}

var ErrThereAreNotEnoughSymbolsToWriteOff = fmt.Errorf("there are not enough symbols to write off")

func (b *translationService) Translate(userID uuid.UUID, text string) (uuid.UUID, error) {
	res, err := b.balanceService.CanWriteOf(userID, len(text))
	if err != nil {
		return uuid.UUID{}, err
	}
	if !res {
		return uuid.UUID{}, ErrThereAreNotEnoughSymbolsToWriteOff
	}
	var translationID domain.TranslationID
	err = b.unitOfWorkFactory.NewUnitOfWork(func(provider RepositoryProvider) error {
		translationID, err = domain.NewTranslationManager(provider.TranslationRepo()).AddTranslation(text, userID)
		if err != nil {
			return err
		}
		return nil
	})
	//TODO: внести добавление в очередь в транзакцию, либо проставлять ошибочный статус при ошиьке добавления в очередь
	b.translationQueue.AddTask(Task{
		TranslationID: uuid.UUID(translationID),
		Text:          text,
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

func NewTranslationService(unitOfWorkFactory UnitOfWorkFactory, translationQueue TranslationQueue, balanceService BalanceService, translationQueryService TranslationQueryService) TranslationService {
	return &translationService{
		unitOfWorkFactory:       unitOfWorkFactory,
		translationQueue:        translationQueue,
		balanceService:          balanceService,
		translationQueryService: translationQueryService,
	}
}
