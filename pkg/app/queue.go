package app

import (
	"github.com/google/uuid"
)

type TranslationQueue interface {
	AddTask(task Task)
	Start()
}

type Task struct {
	TranslationID uuid.UUID
	Text          string
}
