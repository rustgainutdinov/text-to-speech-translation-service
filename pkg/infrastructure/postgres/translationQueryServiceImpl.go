package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/app"
)

type translationQueryServiceImpl struct {
	db pg.DBI
}

func (t *translationQueryServiceImpl) GetTranslationData(translationID uuid.UUID) (app.TranslationDTO, error) {
	var translations []sqlxTranslationStatusAndData
	_, err := t.db.Query(&translations, "SELECT status, translated_data, id_user, text_to_translate FROM translation WHERE id_translation=? LIMIT 1", translationID.String())
	if err != nil {
		return nil, err
	}
	if len(translations) == 0 {
		return nil, app.ErrTranslationIsNotFound
	}
	return &translation{
		status:          translations[0].Status,
		translatedData:  translations[0].TranslatedData,
		idUser:          translations[0].IDUser,
		textToTranslate: translations[0].TextToTranslate,
	}, nil
}

func NewTranslationQueryService(db pg.DBI) app.TranslationQueryService {
	return &translationQueryServiceImpl{db: db}
}

type translation struct {
	status          int
	translatedData  string
	idUser          string
	textToTranslate string
}

func (t *translation) Status() int {
	return t.status
}

func (t *translation) TranslatedData() string {
	return t.translatedData
}

func (t *translation) UserID() string {
	return t.idUser
}

func (t *translation) Text() string {
	return t.textToTranslate
}

type sqlxTranslationStatusAndData struct {
	Status          int    `db:"status"`
	TranslatedData  string `db:"translated_data"`
	IDUser          string `db:"id_user"`
	TextToTranslate string `db:"text_to_translate"`
}
