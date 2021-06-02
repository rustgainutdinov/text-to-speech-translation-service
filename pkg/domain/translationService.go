package domain

import (
	"fmt"
	"github.com/google/uuid"
)

type TranslationService interface {
	AddTranslation(text string, userID uuid.UUID) (TranslationID, error)
	Translate(id TranslationID) error
}

var ErrThereAreNotEnoughSymbolsToWriteOff = fmt.Errorf("there are not enough symbols to write off")

type translationService struct {
	translationQueue            TranslationQueue
	textToSpeechService         TextToSpeechService
	translationTextToSpeechRepo TranslationRepo
	balanceService              BalanceService
}

func (t *translationService) AddTranslation(text string, userID uuid.UUID) (TranslationID, error) {
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

func (t *translationService) Translate(id TranslationID) error {
	translation, err := t.translationTextToSpeechRepo.FindOne(id)
	if err != nil {
		return err
	}
	translation.SpeechData, err = t.textToSpeechService.Translate(translation.Text)
	translation.Status = TranslationStatusSuccess
	if err != nil {
		translation.Status = TranslationStatusError
	}
	err2 := t.translationTextToSpeechRepo.Store(translation)
	if err != nil {
		return err
	}
	return err2
}

func NewTranslationService(translationQueue TranslationQueue, translationTextToSpeechRepo TranslationRepo, textToSpeechService TextToSpeechService, balanceService BalanceService) TranslationService {
	return &translationService{
		translationQueue:            translationQueue,
		translationTextToSpeechRepo: translationTextToSpeechRepo,
		textToSpeechService:         textToSpeechService,
		balanceService:              balanceService,
	}
}
