package infrastructure

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"text-to-speech-translation-service/pkg/app"
)

type translationQueryServiceImpl struct {
	db *sqlx.DB
}

func (t *translationQueryServiceImpl) GetTranslationData(translationID uuid.UUID) (app.TranslationDTO, error) {
	var translations []sqlxTranslationStatusAndData
	err := t.db.Select(&translations, "SELECT status, translated_data FROM translation WHERE id_translation=$1 LIMIT 1", translationID.String())
	if err != nil {
		return nil, err
	}
	if len(translations) == 0 {
		return nil, app.ErrTranslationIsNotFound
	}
	return &translation{
		status:         translations[0].Status,
		translatedData: translations[0].TranslatedData,
	}, nil
}

func NewTranslationQueryService(db *sqlx.DB) app.TranslationQueryService {
	return &translationQueryServiceImpl{db: db}
}

type translation struct {
	status         int
	translatedData string
}

func (t *translation) Status() int {
	return t.status
}

func (t *translation) TranslatedData() string {
	return t.translatedData
}

type sqlxTranslationStatusAndData struct {
	Status         int    `db:"status"`
	TranslatedData string `db:"translated_data"`
}
