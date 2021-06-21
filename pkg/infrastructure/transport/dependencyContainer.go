package transport

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/streadway/amqp"
	"text-to-speech-translation-service/pkg/app"
	"text-to-speech-translation-service/pkg/domain"
	"text-to-speech-translation-service/pkg/infrastructure/externalServices"
	"text-to-speech-translation-service/pkg/infrastructure/postgres"
	"text-to-speech-translation-service/pkg/infrastructure/queue"
)

type DependencyContainer interface {
	newAppTranslationService() app.TranslationService
	newTranslationQueue() domain.TranslationQueue
	newBalanceService() domain.BalanceService
	newExternalTextToSpeechService() domain.ExternalTextToSpeech
	newTranslationQueryService() app.TranslationQueryService
	newExternalEventBroker() domain.ExternalEventBroker
}

type dependencyContainer struct {
	db                pg.DBI
	envConf           Config
	translationQueue  domain.TranslationQueue
	rabbitMqChannel   *amqp.Channel
	unitOfWorkFactory domain.UnitOfWorkFactory
}

func (d *dependencyContainer) newAppTranslationService() app.TranslationService {
	return app.NewTranslationService(d.unitOfWorkFactory, d.newTranslationQueue(), d.newBalanceService(), d.newTranslationQueryService())
}

func (d *dependencyContainer) newTranslationQueue() domain.TranslationQueue {
	return d.translationQueue
}

func (d *dependencyContainer) newBalanceService() domain.BalanceService {
	return externalServices.NewBalanceService(d.envConf.BalanceServiceAddress)
}

func (d *dependencyContainer) newExternalTextToSpeechService() domain.ExternalTextToSpeech {
	return externalServices.NewExternalTextToSpeechService()
}

func (d *dependencyContainer) newTranslationQueryService() app.TranslationQueryService {
	return postgres.NewTranslationQueryService(d.db)
}

func (d *dependencyContainer) newExternalEventBroker() domain.ExternalEventBroker {
	q, err := d.rabbitMqChannel.QueueDeclare("textTranslated", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return externalServices.NewExternalEventBroker(d.rabbitMqChannel, &q)
}

func NewDependencyContainer(db pg.DBI, envConf Config, rabbitMqChannel *amqp.Channel) DependencyContainer {
	d := dependencyContainer{db: db, envConf: envConf, rabbitMqChannel: rabbitMqChannel, unitOfWorkFactory: &postgres.UnitOfWorkFactory{Client: db}}
	d.translationQueue = queue.NewQueue(domain.NewTextToSpeechService(d.unitOfWorkFactory, d.newExternalTextToSpeechService(), d.newExternalEventBroker()))
	return &d
}
