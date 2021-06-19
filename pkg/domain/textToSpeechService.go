package domain

type TextToSpeechService interface {
	Translate(id TranslationID) error
}

type textToSpeechService struct {
	translationTextToSpeechRepo TranslationRepo
	externalTextToSpeech        ExternalTextToSpeech
	externalEventBroker         ExternalEventBroker
}

func (t *textToSpeechService) Translate(id TranslationID) error {
	translation, err := t.translationTextToSpeechRepo.FindOne(id)
	if err != nil {
		return err
	}
	translation.SpeechData, err = t.externalTextToSpeech.Translate(translation.Text)
	translation.Status = TranslationStatusSuccess
	if err != nil {
		translation.Status = TranslationStatusError
	}
	err = t.translationTextToSpeechRepo.Store(translation)
	if err != nil {
		return err
	}
	return t.externalEventBroker.TextTranslated(translation.UserID, len(translation.Text))
}

func NewTextToSpeechService(translationTextToSpeechRepo TranslationRepo, externalTextToSpeech ExternalTextToSpeech, externalEventBroker ExternalEventBroker) TextToSpeechService {
	return &textToSpeechService{
		translationTextToSpeechRepo: translationTextToSpeechRepo,
		externalTextToSpeech:        externalTextToSpeech,
		externalEventBroker:         externalEventBroker,
	}
}
