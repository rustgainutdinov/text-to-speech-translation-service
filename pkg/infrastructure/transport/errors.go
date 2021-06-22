package transport

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"text-to-speech-translation-service/pkg/app/dataProvider"
	"text-to-speech-translation-service/pkg/app/service"
	"text-to-speech-translation-service/pkg/domain"
)

func TranslateError(err error) error {
	if _, ok := status.FromError(err); ok {
		return err
	}
	if err == domain.ErrTranslationIsNotFound || err == dataProvider.ErrTranslationIsNotFound {
		return status.Errorf(codes.NotFound, err.Error())
	}
	if err == service.ErrThereAreNotEnoughSymbolsToWriteOff {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	return status.Errorf(codes.Internal, err.Error())
}
