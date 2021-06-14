package infrastructure

import (
	"context"
	"text-to-speech-translation-service/api"
)

type TranslationServer struct{}

func (t *TranslationServer) Translate(_ context.Context, req *api.TranslationRequest) (*api.TranslationID, error) {
	return &api.TranslationID{
		TranslationID: "228",
	}, nil
}

func (t *TranslationServer) GetTranslationStatus(_ context.Context, req *api.TranslationID) (*api.TranslationStatus, error) {
	return &api.TranslationStatus{TranslationStatus: api.TranslationStatusEnum_SUCCESS}, nil
}

func (t *TranslationServer) GetTranslationData(_ context.Context, req *api.TranslationID) (*api.TranslationData, error) {
	return &api.TranslationData{Text: "some text"}, nil
}
