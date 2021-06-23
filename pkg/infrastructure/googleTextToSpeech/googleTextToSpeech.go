package googleTextToSpeech

import (
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	"encoding/base64"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"text-to-speech-translation-service/pkg/app/service"
)

type googleTextToSpeech struct {
	client  *texttospeech.Client
	context context.Context
}

func (t *googleTextToSpeech) Translate(text string) (string, error) {
	err := t.tryToInitClient()
	if err != nil {
		return "", err
	}
	req := generateRequest(text)
	resp, err := t.client.SynthesizeSpeech(t.context, &req)
	if err != nil {
		return "", err
	}
	sEnc := base64.StdEncoding.EncodeToString(resp.AudioContent)
	return sEnc, nil
}

func (t *googleTextToSpeech) tryToInitClient() error {
	if t.client == nil {
		ctx := context.Background()
		client, err := texttospeech.NewClient(ctx)
		if err != nil {
			return err
		}
		t.client = client
		t.context = ctx
	}
	return nil
}

func generateRequest(text string) texttospeechpb.SynthesizeSpeechRequest {
	return texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "ru-RU",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}
}

func NewGoogleTextToSpeechService() service.ExternalTextToSpeech {
	return &googleTextToSpeech{}
}
