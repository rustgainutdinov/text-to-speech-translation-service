package infrastructure

import "text-to-speech-translation-service/pkg/domain"

type externalTextToSpeechImpl struct{}

func (t *externalTextToSpeechImpl) Translate(text string) (string, error) {
	return "olala", nil
}

func NewExternalTextToSpeechService() domain.ExternalTextToSpeech {
	return &externalTextToSpeechImpl{}
}
