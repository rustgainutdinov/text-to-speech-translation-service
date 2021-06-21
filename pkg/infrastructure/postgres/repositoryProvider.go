package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"text-to-speech-translation-service/pkg/domain"
)

type unitOfWork struct {
	tx *pg.Tx
}

func (u *unitOfWork) TranslationRepo() domain.TranslationRepo {
	return NewTranslationRepo(u.tx)
}

func (u *unitOfWork) Complete(err error) error {
	if err != nil {
		err2 := u.tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	return u.tx.Commit()
}

type UnitOfWorkFactory struct {
	Client pg.DBI
}

func (u *UnitOfWorkFactory) NewUnitOfWork(f func(provider domain.RepositoryProvider) error) error {
	return u.Client.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		unitOfWork := &unitOfWork{tx: tx}
		return f(unitOfWork)
	})
}
