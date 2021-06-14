package domain

type ExternalTextToSpeech interface {
	Translate(text string) (string, error)
}
