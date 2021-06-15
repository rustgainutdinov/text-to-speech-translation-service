package infrastructure

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"text-to-speech-translation-service/pkg/domain"
)

type translationRepo struct {
	db *sqlx.DB
}

func (t *translationRepo) Store(translation domain.Translation) error {
	_, err := t.db.Query(
		`INSERT INTO translation (id_translation, id_user, status, text_to_translate, translated_data)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id_translation)
			DO UPDATE SET id_user           = $2,
						  status            = $3,
						  text_to_translate = $4,
						  translated_data   = $5;`,
		uuid.UUID(translation.ID).String(), translation.UserID, translation.Status, translation.Text, translation.SpeechData)
	return err
}

func (t *translationRepo) FindOne(translationID domain.TranslationID) (domain.Translation, error) {
	var translations []sqlxTranslation
	err := t.db.Select(&translations, "SELECT * FROM translation WHERE id_translation=$1 LIMIT 1", uuid.UUID(translationID).String())
	if err != nil {
		return domain.Translation{}, err
	}
	if len(translations) == 0 {
		return domain.Translation{}, domain.ErrTranslationIsNotFound
	}
	translationUUID, err := uuid.Parse(translations[0].ID)
	if err != nil {
		return domain.Translation{}, err
	}
	userUUID, err := uuid.Parse(translations[0].UserID)
	if err != nil {
		return domain.Translation{}, err
	}
	return domain.Translation{
		ID:         domain.TranslationID(translationUUID),
		UserID:     userUUID,
		Text:       translations[0].TextToTranslate,
		Status:     translations[0].Status,
		SpeechData: translations[0].TranslatedData,
	}, nil
}

func NewTranslationRepo(db *sqlx.DB) domain.TranslationRepo {
	return &translationRepo{db: db}
}

type sqlxTranslation struct {
	ID              string `db:"id_translation"`
	UserID          string `db:"id_user"`
	Status          int    `db:"status"`
	TextToTranslate string `db:"text_to_translate"`
	TranslatedData  string `db:"translated_data"`
}
