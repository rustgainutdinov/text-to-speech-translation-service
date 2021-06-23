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
	rabbitMqMessage := rabbitMqMessage{Type: 0, Data: textTranslatedInfo{UserID: userID, Score: score}}
	body, err := json.Marshal(rabbitMqMessage)
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

type rabbitMqMessage struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

type textTranslatedInfo struct {
	UserID string `json:"userID"`
	Score  int    `json:"score"`
}
