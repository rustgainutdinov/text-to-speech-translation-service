package textToSpeech

import (
	"text-to-speech-translation-service/pkg/app/externalService/textToSpeech"
)

type externalTextToSpeech struct{}

func (t *externalTextToSpeech) Translate(text string) (string, error) {
	return text + " - translated", nil
}

func NewExternalTextToSpeechService() textToSpeech.ExternalTextToSpeech {
	return &externalTextToSpeech{}
}
