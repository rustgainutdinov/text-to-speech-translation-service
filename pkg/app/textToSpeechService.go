package app

import (
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/domain"
)

type TextToSpeechService interface {
	Translate(id uuid.UUID) error
}

type textToSpeechService struct {
	unitOfWorkFactory       UnitOfWorkFactory
	externalTextToSpeech    ExternalTextToSpeech
	externalEventBroker     ExternalEventBroker
	translationQueryService TranslationQueryService
}

func (t *textToSpeechService) Translate(id uuid.UUID) error {
	translation, err := t.translationQueryService.GetTranslationData(id)
	if err != nil {
		return err
	}
	translatedData, err := t.externalTextToSpeech.Translate(translation.Text())
	return t.unitOfWorkFactory.NewUnitOfWork(func(provider RepositoryProvider) error {
		err = domain.NewTranslationManager(provider.TranslationRepo()).SaveTranslatedData(domain.TranslationID(id), translatedData)
		if err != nil {
			return err
		}
		return t.externalEventBroker.TextTranslated(translation.UserID(), len(translation.Text()))
	})
}

func NewTextToSpeechService(unitOfWorkFactory UnitOfWorkFactory, externalTextToSpeech ExternalTextToSpeech, externalEventBroker ExternalEventBroker, translationQueryService TranslationQueryService) TextToSpeechService {
	return &textToSpeechService{
		unitOfWorkFactory:       unitOfWorkFactory,
		externalTextToSpeech:    externalTextToSpeech,
		externalEventBroker:     externalEventBroker,
		translationQueryService: translationQueryService,
	}
}
