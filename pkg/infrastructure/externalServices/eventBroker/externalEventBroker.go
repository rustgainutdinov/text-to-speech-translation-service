package eventBroker

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	externalEventBroker2 "text-to-speech-translation-service/pkg/app/externalService/eventBroker"
)

type externalEventBroker struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

func (e *externalEventBroker) TextTranslated(userID string, score int) error {
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

func NewExternalEventBroker(channel *amqp.Channel, queue *amqp.Queue) externalEventBroker2.ExternalEventBroker {
	return &externalEventBroker{channel: channel, queue: queue}
}

type textTranslatedInfo struct {
	UserID string `json:"userID"`
	Score  int    `json:"score"`
}
