package domain

import "github.com/google/uuid"

type TranslationService interface {
	AddTextToSpeechTranslation(text string) (uuid.UUID, error)
	TranslateTextToSpeech(id TranslationTextToSpeechID) error
}

type translationService struct {
	translationQueue            TranslationQueue
	textToSpeechService         TextToSpeechService
	translationTextToSpeechRepo TranslationTextToSpeechRepo
}

func (t *translationService) AddTextToSpeechTranslation(text string) (uuid.UUID, error) {
	translationID := uuid.New()
	translation := TranslationTextToSpeech{
		ID:     TranslationTextToSpeechID(translationID),
		Text:   text,
		Status: TranslationStatusWaiting,
	}
	err := t.translationTextToSpeechRepo.Store(translation)
	if err != nil {
		return uuid.UUID{}, err
	}
	t.translationQueue.AddJob(Task{
		TaskType: TaskTypeSpeechToTextTranslation,
		ID:       TaskID(uuid.New()),
		Data:     nil,
	})
	return translationID, nil
}

func (t *translationService) TranslateTextToSpeech(id TranslationTextToSpeechID) error {
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

func NewTranslationService(translationQueue TranslationQueue, translationTextToSpeechRepo TranslationTextToSpeechRepo, textToSpeechService TextToSpeechService) TranslationService {
	return &translationService{
		translationQueue:            translationQueue,
		translationTextToSpeechRepo: translationTextToSpeechRepo,
		textToSpeechService:         textToSpeechService,
	}
}
