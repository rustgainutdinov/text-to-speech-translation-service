package query

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/app/dataProvider"
)

type translationQueryServiceImpl struct {
	db pg.DBI
}

func (t *translationQueryServiceImpl) GetTranslationData(translationID uuid.UUID) (dataProvider.TranslationDTO, error) {
	var translationData sqlxTranslationStatusAndData
	_, err := t.db.QueryOne(&translationData, "SELECT status, translated_data, id_user, text_to_translate FROM translation WHERE id_translation=?", translationID.String())
	if err != nil {
		return nil, err
	}
	//TODO: добавить обработку ошибки not found (ErrTranslationIsNotFound)
	return &translation{
		status:          translationData.Status,
		translatedData:  translationData.TranslatedData,
		idUser:          translationData.IDUser,
		textToTranslate: translationData.TextToTranslate,
	}, nil
}

func NewTranslationQueryService(db pg.DBI) dataProvider.TranslationQueryService {
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
