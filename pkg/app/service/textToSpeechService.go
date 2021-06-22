package service

import (
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/app/dataProvider"
	externalEventBroker2 "text-to-speech-translation-service/pkg/app/externalService/eventBroker"
	"text-to-speech-translation-service/pkg/app/externalService/textToSpeech"
	"text-to-speech-translation-service/pkg/domain"
)

type TextToSpeechService interface {
	Translate(id uuid.UUID) error
}

type textToSpeechService struct {
	unitOfWorkFactory       dataProvider.UnitOfWorkFactory
	externalTextToSpeech    textToSpeech.ExternalTextToSpeech
	externalEventBroker     externalEventBroker2.ExternalEventBroker
	translationQueryService dataProvider.TranslationQueryService
}

func (t *textToSpeechService) Translate(id uuid.UUID) error {
	translation, err := t.translationQueryService.GetTranslationData(id)
	if err != nil {
		return err
	}
	translatedData, err := t.externalTextToSpeech.Translate(translation.Text())
	err = t.unitOfWorkFactory.NewUnitOfWork(func(provider dataProvider.RepositoryProvider) error {
		err = domain.NewTranslationManager(provider.TranslationRepo()).SaveTranslation(domain.TranslationID(id), translatedData)
		if err != nil {
			return err
		}
		return t.externalEventBroker.TextTranslated(translation.UserID(), len(translation.Text()))
	})
	if err == nil {
		return nil
	}
	err2 := t.unitOfWorkFactory.NewUnitOfWork(func(provider dataProvider.RepositoryProvider) error {
		return domain.NewTranslationManager(provider.TranslationRepo()).MarkTranslationAsErrored(domain.TranslationID(id))
	})
	if err2 != nil {
		return err2
	}
	return err
}

func NewTextToSpeechService(unitOfWorkFactory dataProvider.UnitOfWorkFactory, externalTextToSpeech textToSpeech.ExternalTextToSpeech, externalEventBroker externalEventBroker2.ExternalEventBroker, translationQueryService dataProvider.TranslationQueryService) TextToSpeechService {
	return &textToSpeechService{
		unitOfWorkFactory:       unitOfWorkFactory,
		externalTextToSpeech:    externalTextToSpeech,
		externalEventBroker:     externalEventBroker,
		translationQueryService: translationQueryService,
	}
}
