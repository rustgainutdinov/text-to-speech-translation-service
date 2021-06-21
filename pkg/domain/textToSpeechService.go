package domain

type TextToSpeechService interface {
	Translate(id TranslationID) error
}

type textToSpeechService struct {
	unitOfWorkFactory    UnitOfWorkFactory
	externalTextToSpeech ExternalTextToSpeech
	externalEventBroker  ExternalEventBroker
}

func (t *textToSpeechService) Translate(id TranslationID) error {
	return t.unitOfWorkFactory.NewUnitOfWork(func(provider RepositoryProvider) error {
		translationRepo := provider.TranslationRepo()
		translation, err := translationRepo.FindOne(id)
		if err != nil {
			return err
		}
		translation.SpeechData, err = t.externalTextToSpeech.Translate(translation.Text)
		translation.Status = TranslationStatusSuccess
		if err != nil {
			translation.Status = TranslationStatusError
		}
		err = translationRepo.Store(translation)
		if err != nil {
			return err
		}
		return t.externalEventBroker.TextTranslated(translation.UserID, len(translation.Text))
	})
}

func NewTextToSpeechService(unitOfWorkFactory UnitOfWorkFactory, externalTextToSpeech ExternalTextToSpeech, externalEventBroker ExternalEventBroker) TextToSpeechService {
	return &textToSpeechService{
		unitOfWorkFactory:    unitOfWorkFactory,
		externalTextToSpeech: externalTextToSpeech,
		externalEventBroker:  externalEventBroker,
	}
}
