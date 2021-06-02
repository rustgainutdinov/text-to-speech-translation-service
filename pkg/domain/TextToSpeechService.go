package domain

type TextToSpeechService interface {
	Translate(text string) (string, error)
}
