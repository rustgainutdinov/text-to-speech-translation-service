package infrastructure

import (
	"context"
	"github.com/google/uuid"
	"text-to-speech-translation-service/api"
	"text-to-speech-translation-service/pkg/domain"
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
	translationID, err := uuid.Parse(req.TranslationID)
	if err != nil {
		return nil, err
	}
	status, err := t.DependencyContainer.newAppTranslationService().GetTranslationStatus(translationID)
	if err != nil {
		return nil, err
	}
	return &api.TranslationStatus{TranslationStatus: convertDomainStatusToAPI(status)}, err
}

func (t *TranslationServer) GetTranslationData(_ context.Context, req *api.TranslationID) (*api.TranslationData, error) {
	translationID, err := uuid.Parse(req.TranslationID)
	if err != nil {
		return nil, err
	}
	translatedData, err := t.DependencyContainer.newAppTranslationService().GetTranslationData(translationID)
	if err != nil {
		return nil, err
	}
	return &api.TranslationData{Text: translatedData}, nil
}

func convertDomainStatusToAPI(domainStatus int) api.TranslationStatusEnum {
	switch domainStatus {
	case domain.TranslationStatusWaiting:
		return api.TranslationStatusEnum_WAITING
	case domain.TranslationStatusSuccess:
		return api.TranslationStatusEnum_SUCCESS
	case domain.TranslationStatusError:
		return api.TranslationStatusEnum_ERROR
	default:
		return api.TranslationStatusEnum_ERROR
	}
}
