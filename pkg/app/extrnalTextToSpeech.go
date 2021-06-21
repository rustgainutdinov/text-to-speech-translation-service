package app

type ExternalTextToSpeech interface {
	Translate(text string) (string, error)
}
