package externalServices

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"text-to-speech-translation-service/pkg/app"
)

type ExternalEventBrokerImpl struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

func (e *ExternalEventBrokerImpl) TextTranslated(userID string, score int) error {
	translatedInfo := textTranslatedInfo{UserID: userID, Score: score}
	body, err := json.Marshal(translatedInfo)
	if err != nil {
		return err
	}
	err = e.channel.Publish("", e.queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		return err
	}
	log.Printf("Text translated event: userID-%s score-%d", userID, score)
	return nil
}

func NewExternalEventBroker(channel *amqp.Channel, queue *amqp.Queue) app.ExternalEventBroker {
	return &ExternalEventBrokerImpl{channel: channel, queue: queue}
}

type textTranslatedInfo struct {
	UserID string `json:"userID"`
	Score  int    `json:"score"`
}
