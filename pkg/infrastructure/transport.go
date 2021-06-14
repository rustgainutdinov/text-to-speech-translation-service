package infrastructure

import (
	"context"
	"github.com/google/uuid"
	"text-to-speech-translation-service/api"
)

type TranslationServer struct {
	DependencyContainer DependencyContainer
}

func (t *TranslationServer) Translate(_ context.Context, req *api.TranslationRequest) (*api.TranslationID, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}
	translationID, err := t.DependencyContainer.newAppTranslationService().Translate(userID, req.Text)
	if err != nil {
		return nil, err
	}
	return &api.TranslationID{TranslationID: translationID.String()}, nil
}

func (t *TranslationServer) GetTranslationStatus(_ context.Context, req *api.TranslationID) (*api.TranslationStatus, error) {
	return &api.TranslationStatus{TranslationStatus: api.TranslationStatusEnum_SUCCESS}, nil
}

func (t *TranslationServer) GetTranslationData(_ context.Context, req *api.TranslationID) (*api.TranslationData, error) {
	return &api.TranslationData{Text: "some text"}, nil
}
