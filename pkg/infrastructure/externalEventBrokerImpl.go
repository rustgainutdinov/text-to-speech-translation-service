package infrastructure

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
	"text-to-speech-translation-service/pkg/domain"
)

type ExternalEventBrokerImpl struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

func (e *ExternalEventBrokerImpl) TextTranslated(userID uuid.UUID, amountOfSymbols int) error {
	translatedInfo := textTranslatedInfo{UserID: userID.String(), AmountOfSymbols: amountOfSymbols}
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
	log.Printf("Text translated event: userID-%s amountOfSymbols-%d", userID.String(), amountOfSymbols)
	return nil
}

func NewExternalEventBroker(channel *amqp.Channel, queue *amqp.Queue) domain.ExternalEventBroker {
	return &ExternalEventBrokerImpl{channel: channel, queue: queue}
}

type textTranslatedInfo struct {
	UserID          string `json:"userID"`
	AmountOfSymbols int    `json:"amountOfSymbols"`
}
