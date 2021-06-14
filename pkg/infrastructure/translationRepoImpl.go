package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"text-to-speech-translation-service/pkg/domain"
)

type translationRepo struct {
	db *sqlx.DB
}

func (t *translationRepo) Store(translation domain.Translation) error {
	_, err := t.db.Query(
		`INSERT INTO transaction (id_transaction, id_user, status, text_to_translate, translated_data)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id_transaction)
			DO UPDATE SET id_user           = $2,
						  status            = $3,
						  text_to_translate = $4,
						  translated_data   = $5;`,
		translation.ID, translation.UserID, translation.Status, translation.Text, translation.SpeechData)
	return err
}

func (t *translationRepo) FindOne(translationID domain.TranslationID) (domain.Translation, error) {
	return domain.Translation{}, nil
}

func NewTranslationRepo(db *sqlx.DB) domain.TranslationRepo {
	return &translationRepo{db: db}
}
