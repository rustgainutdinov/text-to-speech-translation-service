package transport

import (
	"context"
	"github.com/google/uuid"
	"text-to-speech-translation-service/api"
	"text-to-speech-translation-service/pkg/domain"
	"text-to-speech-translation-service/pkg/infrastructure"
)

type TranslationServer struct {
	DependencyContainer infrastructure.DependencyContainer
}

func (t *TranslationServer) AddTextToTranslate(_ context.Context, req *api.TranslationRequest) (*api.TranslationID, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}
	translationID, err := t.DependencyContainer.NewTranslationService().AddTextToTranslate(userID, req.Text)
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
	status, err := t.DependencyContainer.NewTranslationService().GetTranslationStatus(translationID)
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
	translatedData, err := t.DependencyContainer.NewTranslationService().GetTranslationData(translationID)
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
