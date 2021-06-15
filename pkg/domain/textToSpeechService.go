package domain

type TextToSpeechService interface {
	Translate(id TranslationID) error
}

type textToSpeechService struct {
	translationTextToSpeechRepo TranslationRepo
	externalTextToSpeech        ExternalTextToSpeech
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
	err2 := t.translationTextToSpeechRepo.Store(translation)
	if err != nil {
		return err
	}
	return err2
}

func NewTextToSpeechService(translationTextToSpeechRepo TranslationRepo, externalTextToSpeech ExternalTextToSpeech) TextToSpeechService {
	return &textToSpeechService{
		translationTextToSpeechRepo: translationTextToSpeechRepo,
		externalTextToSpeech:        externalTextToSpeech,
	}
}
