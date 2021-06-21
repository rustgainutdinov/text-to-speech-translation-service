package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/domain"
)

type translationRepo struct {
	tx *pg.Tx
}

func (t *translationRepo) Store(translation domain.Translation) error {
	_, err := t.tx.Exec(
		`INSERT INTO translation (id_translation, id_user, status, text_to_translate, translated_data)
		VALUES (?id, ?user_id, ?status, ?text_to_translate, ?translated_data)
		ON CONFLICT (id_translation)
			DO UPDATE SET id_user           = ?user_id,
						  status            = ?status,
						  text_to_translate = ?text_to_translate,
						  translated_data   = ?translated_data;`,
		translationDTO{uuid.UUID(translation.ID).String(), translation.UserID.String(), translation.Status, translation.Text, translation.SpeechData},
	)
	return err
}

func (t *translationRepo) FindOne(translationID domain.TranslationID) (domain.Translation, error) {
	var translations []sqlxTranslation
	_, err := t.tx.Query(&translations, "SELECT * FROM translation WHERE id_translation=? LIMIT 1", uuid.UUID(translationID).String())
	if err != nil {
		return domain.Translation{}, err
	}
	if len(translations) == 0 {
		return domain.Translation{}, domain.ErrTranslationIsNotFound
	}
	translationUUID, err := uuid.Parse(translations[0].IDTranslation)
	if err != nil {
		return domain.Translation{}, err
	}
	userUUID, err := uuid.Parse(translations[0].IDUser)
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

func NewTranslationRepo(tx *pg.Tx) domain.TranslationRepo {
	return &translationRepo{tx: tx}
}

type translationDTO struct {
	ID              string
	UserID          string
	Status          int
	TextToTranslate string
	TranslatedData  string
}

type sqlxTranslation struct {
	IDTranslation   string `db:"id_translation"`
	IDUser          string `db:"id_user"`
	Status          int    `db:"status"`
	TextToTranslate string `db:"text_to_translate"`
	TranslatedData  string `db:"translated_data"`
}
