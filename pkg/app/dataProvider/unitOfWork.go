package dataProvider

import "text-to-speech-translation-service/pkg/domain"

type RepositoryProvider interface {
	TranslationRepo() domain.TranslationRepo
}

type UnitOfWork interface {
	RepositoryProvider
	Complete(err error) error
}

type UnitOfWorkFactory interface {
	NewUnitOfWork(f func(provider RepositoryProvider) error) error
}
