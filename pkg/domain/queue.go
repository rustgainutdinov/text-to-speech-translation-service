package domain

import (
	"github.com/google/uuid"
)

const (
	TaskTypeSpeechToTextTranslation = iota
)

type TranslationQueue interface {
	AddJob(task Task)
}

type TaskID uuid.UUID

type Task struct {
	TaskType int
	ID       TaskID
	Data     map[string]string
}
