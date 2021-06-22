package infrastructure

import (
	"github.com/go-pg/pg/v10"
	"github.com/streadway/amqp"
	"text-to-speech-translation-service/pkg/app"
	query2 "text-to-speech-translation-service/pkg/app/dataProvider"
	balanceService2 "text-to-speech-translation-service/pkg/app/externalService/balanceService"
	externalEventBroker2 "text-to-speech-translation-service/pkg/app/externalService/eventBroker"
	textToSpeech2 "text-to-speech-translation-service/pkg/app/externalService/textToSpeech"
	"text-to-speech-translation-service/pkg/app/service"
	"text-to-speech-translation-service/pkg/infrastructure/externalServices/balanceService"
	"text-to-speech-translation-service/pkg/infrastructure/externalServices/eventBroker"
	"text-to-speech-translation-service/pkg/infrastructure/externalServices/textToSpeech"
	"text-to-speech-translation-service/pkg/infrastructure/postgres/query"
	"text-to-speech-translation-service/pkg/infrastructure/postgres/repo"
	"text-to-speech-translation-service/pkg/infrastructure/queue"
)

type DependencyContainer interface {
	NewTranslationService() service.TranslationService
}

type dependencyContainer struct {
	db                 pg.DBI
	envConf            Config
	translationQueue   app.Queue
	rabbitMqChannel    *amqp.Channel
	unitOfWorkFactory  query2.UnitOfWorkFactory
	translationService service.TranslationService
}

func (d *dependencyContainer) NewTranslationService() service.TranslationService {
	if d.translationService == nil {
		d.translationService = service.NewTranslationService(d.unitOfWorkFactory, d.newTranslationQueue(), d.newBalanceService(), d.newTranslationQueryService())
	}
	return d.translationService
}

func (d *dependencyContainer) newTranslationQueue() app.Queue {
	return d.translationQueue
}

func (d *dependencyContainer) newBalanceService() balanceService2.BalanceService {
	return balanceService.NewBalanceService(d.envConf.BalanceServiceAddress)
}

func (d *dependencyContainer) newExternalTextToSpeechService() textToSpeech2.ExternalTextToSpeech {
	return textToSpeech.NewExternalTextToSpeechService()
}

func (d *dependencyContainer) newTranslationQueryService() query2.TranslationQueryService {
	return query.NewTranslationQueryService(d.db)
}

func (d *dependencyContainer) newExternalEventBroker() (externalEventBroker2.ExternalEventBroker, error) {
	q, err := d.rabbitMqChannel.QueueDeclare("textTranslated", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return eventBroker.NewExternalEventBroker(d.rabbitMqChannel, &q), nil
}

func NewDependencyContainer(db pg.DBI, envConf Config, rabbitMqChannel *amqp.Channel) (DependencyContainer, error) {
	d := dependencyContainer{db: db, envConf: envConf, rabbitMqChannel: rabbitMqChannel, unitOfWorkFactory: &repo.UnitOfWorkFactory{Client: db}}
	newExternalEventBroker, err := d.newExternalEventBroker()
	if err != nil {
		return nil, err
	}
	d.translationQueue = queue.NewQueue(service.NewTextToSpeechService(d.unitOfWorkFactory, d.newExternalTextToSpeechService(), newExternalEventBroker, d.newTranslationQueryService()))
	return &d, nil
}
