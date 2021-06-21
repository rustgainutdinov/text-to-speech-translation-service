package domain

type RepositoryProvider interface {
	TranslationRepo() TranslationRepo
}

type UnitOfWork interface {
	RepositoryProvider
	Complete(err error) error
}

type UnitOfWorkFactory interface {
	NewUnitOfWork(f func(provider RepositoryProvider) error) error
}
