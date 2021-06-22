package eventBroker

type ExternalEventBroker interface {
	TextTranslated(userID string, score int) error
}
