package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"text-to-speech-translation-service/pkg/app"
	"text-to-speech-translation-service/pkg/domain"
)

type DependencyContainer interface {
	newAppTranslationService() app.TranslationService
	newDomainTextToSpeechService() domain.TextToSpeechService
	newTranslationRepo() domain.TranslationRepo
	newTranslationQueue() domain.TranslationQueue
	newBalanceService() domain.BalanceService
	newExternalTextToSpeechService() domain.ExternalTextToSpeech
	newTranslationQueryService() app.TranslationQueryService
}

type dependencyContainer struct {
	db               *sqlx.DB
	envConf          Config
	translationQueue domain.TranslationQueue
}

func (d *dependencyContainer) newAppTranslationService() app.TranslationService {
	return app.NewTranslationService(d.newTranslationRepo(), d.newTranslationQueue(), d.newBalanceService(), d.newTranslationQueryService())
}

func (d *dependencyContainer) newDomainTextToSpeechService() domain.TextToSpeechService {
	return domain.NewTextToSpeechService(d.newTranslationRepo(), d.newExternalTextToSpeechService())
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

func NewDependencyContainer(db *sqlx.DB, envConf Config) DependencyContainer {
	d := dependencyContainer{db: db, envConf: envConf}
	d.translationQueue = NewQueue(domain.NewTextToSpeechService(d.newTranslationRepo(), d.newExternalTextToSpeechService()))
	return &d
}
