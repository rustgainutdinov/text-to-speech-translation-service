package infrastructure

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"text-to-speech-translation-service/pkg/domain"
)

type translationRepo struct {
	db *sqlx.DB
}

func (t *translationRepo) Store(translation domain.Translation) error {
	fmt.Println("translation stored")
	return nil
}

func (t *translationRepo) FindOne(translationID domain.TranslationID) (domain.Translation, error) {
	return domain.Translation{}, nil
}

func NewTranslationRepo(db *sqlx.DB) domain.TranslationRepo {
	return &translationRepo{db: db}
}
