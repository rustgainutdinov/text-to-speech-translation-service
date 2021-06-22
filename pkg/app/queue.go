package app

import (
	"github.com/google/uuid"
)

const (
	TextTranslatedTaskType = iota
)

type Queue interface {
	AddTask(task Task)
	Start()
}

type Task struct {
	Type int
	Data interface{}
}

type TextTranslated struct {
	TranslationID uuid.UUID
	Text          string
}
