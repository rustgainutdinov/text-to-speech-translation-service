package textToSpeech

import (
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	"encoding/base64"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"text-to-speech-translation-service/pkg/app/externalService/textToSpeech"
)

type externalTextToSpeech struct{}

func (t *externalTextToSpeech) Translate(text string) (string, error) {
	ctx := context.Background()
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()
	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}
	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		return "", err
	}
	sEnc := base64.StdEncoding.EncodeToString(resp.AudioContent)
	return sEnc, nil
}

func NewExternalTextToSpeechService() textToSpeech.ExternalTextToSpeech {
	return &externalTextToSpeech{}
}
