package textToSpeech

type ExternalTextToSpeech interface {
	Translate(text string) (string, error)
}
