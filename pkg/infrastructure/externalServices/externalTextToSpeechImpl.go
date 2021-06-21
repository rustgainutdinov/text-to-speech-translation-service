package externalServices

import (
	"text-to-speech-translation-service/pkg/app"
)

type externalTextToSpeechImpl struct{}

func (t *externalTextToSpeechImpl) Translate(text string) (string, error) {
	return "olala", nil
}

func NewExternalTextToSpeechService() app.ExternalTextToSpeech {
	return &externalTextToSpeechImpl{}
}
