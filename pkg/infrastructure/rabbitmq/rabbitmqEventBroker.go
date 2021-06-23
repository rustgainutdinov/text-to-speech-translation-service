package rabbitmq

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"text-to-speech-translation-service/pkg/app/service"
)

type rabbitmqEventBroker struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

func (e *rabbitmqEventBroker) TextTranslated(userID string, score int) error {
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
	log.WithFields(log.Fields{"userID": userID, "score": score}).Info("Text translated")
	return nil
}

func NewRabbitmqEventBroker(channel *amqp.Channel, queue *amqp.Queue) service.ExternalEventBroker {
	return &rabbitmqEventBroker{channel: channel, queue: queue}
}

type textTranslatedInfo struct {
	UserID string `json:"userID"`
	Score  int    `json:"score"`
}
