package infrastructure

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
	"text-to-speech-translation-service/pkg/app"
	"text-to-speech-translation-service/pkg/domain"
)

type DependencyContainer interface {
	newAppTranslationService() app.TranslationService
	newTranslationRepo() domain.TranslationRepo
	newTranslationQueue() domain.TranslationQueue
	newBalanceService() domain.BalanceService
	newExternalTextToSpeechService() domain.ExternalTextToSpeech
	newTranslationQueryService() app.TranslationQueryService
	newExternalEventBroker() domain.ExternalEventBroker
}

type dependencyContainer struct {
	db               *sqlx.DB
	envConf          Config
	translationQueue domain.TranslationQueue
	rabbitMqChannel  *amqp.Channel
}

func (d *dependencyContainer) newAppTranslationService() app.TranslationService {
	return app.NewTranslationService(d.newTranslationRepo(), d.newTranslationQueue(), d.newBalanceService(), d.newTranslationQueryService())
}

func (d *dependencyContainer) newTranslationRepo() domain.TranslationRepo {
	return NewTranslationRepo(d.db)
}

func (d *dependencyContainer) newTranslationQueue() domain.TranslationQueue {
	return d.translationQueue
}

func (d *dependencyContainer) newBalanceService() domain.BalanceService {
	return NewBalanceService(d.envConf.BalanceServiceAddress)
}

func (d *dependencyContainer) newExternalTextToSpeechService() domain.ExternalTextToSpeech {
	return NewExternalTextToSpeechService()
}

func (d *dependencyContainer) newTranslationQueryService() app.TranslationQueryService {
	return NewTranslationQueryService(d.db)
}

func (d *dependencyContainer) newExternalEventBroker() domain.ExternalEventBroker {
	queue, err := d.rabbitMqChannel.QueueDeclare("textTranslated", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return NewExternalEventBroker(d.rabbitMqChannel, &queue)
}

func NewDependencyContainer(db *sqlx.DB, envConf Config, rabbitMqChannel *amqp.Channel) DependencyContainer {
	d := dependencyContainer{db: db, envConf: envConf, rabbitMqChannel: rabbitMqChannel}
	d.translationQueue = NewQueue(domain.NewTextToSpeechService(d.newTranslationRepo(), d.newExternalTextToSpeechService(), d.newExternalEventBroker()))
	return &d
}
