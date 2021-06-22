package repo

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
	var translation sqlxTranslation
	_, err := t.tx.QueryOne(&translation, "SELECT * FROM translation WHERE id_translation=?", uuid.UUID(translationID).String())
	if err != nil {
		if err == pg.ErrNoRows {
			return domain.Translation{}, domain.ErrTranslationIsNotFound
		}
		return domain.Translation{}, err
	}
	translationUUID, err := uuid.Parse(translation.IDTranslation)
	if err != nil {
		return domain.Translation{}, err
	}
	userUUID, err := uuid.Parse(translation.IDUser)
	if err != nil {
		return domain.Translation{}, err
	}
	return domain.Translation{
		ID:         domain.TranslationID(translationUUID),
		UserID:     userUUID,
		Text:       translation.TextToTranslate,
		Status:     translation.Status,
		SpeechData: translation.TranslatedData,
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
