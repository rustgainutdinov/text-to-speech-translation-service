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
	translationQueue            TranslationQueue
	translationTextToSpeechRepo TranslationRepo
	balanceService              BalanceService
}

func (t *translationManager) AddTranslation(text string, userID uuid.UUID) (TranslationID, error) {
	res, err := t.balanceService.CanWriteOf(userID, len(text))
	if err != nil {
		return TranslationID{}, err
	}
	if !res {
		return TranslationID{}, ErrThereAreNotEnoughSymbolsToWriteOff
	}
	translationID := TranslationID(uuid.New())
	translation := Translation{
		ID:     translationID,
		UserID: userID,
		Text:   text,
		Status: TranslationStatusWaiting,
	}
	err = t.translationTextToSpeechRepo.Store(translation)
	if err != nil {
		return TranslationID{}, err
	}
	t.translationQueue.AddTask(Task{
		TranslationID: translationID,
		Text:          text,
	})
	return translationID, nil
}

func NewTranslationManager(translationQueue TranslationQueue, translationRepo TranslationRepo, balanceService BalanceService) TranslationManager {
	return &translationManager{
		translationQueue:            translationQueue,
		translationTextToSpeechRepo: translationRepo,
		balanceService:              balanceService,
	}
}
