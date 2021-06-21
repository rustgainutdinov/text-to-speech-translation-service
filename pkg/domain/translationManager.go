package domain

import (
	"fmt"
	"github.com/google/uuid"
)

type TranslationManager interface {
	AddTranslation(text string, userID uuid.UUID) (TranslationID, error)
}

var ErrThereAreNotEnoughSymbolsToWriteOff = fmt.Errorf("there are not enough symbols to write off")

type translationManager struct {
	translationQueue  TranslationQueue
	unitOfWorkFactory UnitOfWorkFactory
	balanceService    BalanceService
}

func (t *translationManager) AddTranslation(text string, userID uuid.UUID) (TranslationID, error) {
	var translationID TranslationID
	err := t.unitOfWorkFactory.NewUnitOfWork(func(provider RepositoryProvider) error {
		res, err := t.balanceService.CanWriteOf(userID, len(text))
		if err != nil {
			return err
		}
		if !res {
			return ErrThereAreNotEnoughSymbolsToWriteOff
		}
		translationID = TranslationID(uuid.New())
		translation := Translation{
			ID:     translationID,
			UserID: userID,
			Text:   text,
			Status: TranslationStatusWaiting,
		}
		err = provider.TranslationRepo().Store(translation)
		if err != nil {
			return err
		}
		t.translationQueue.AddTask(Task{
			TranslationID: translationID,
			Text:          text,
		})
		return nil
	})
	if err != nil {
		return TranslationID{}, err
	}
	return translationID, nil
}

func NewTranslationManager(translationQueue TranslationQueue, unitOfWorkFactory UnitOfWorkFactory, balanceService BalanceService) TranslationManager {
	return &translationManager{
		translationQueue:  translationQueue,
		unitOfWorkFactory: unitOfWorkFactory,
		balanceService:    balanceService,
	}
}
