package domain

type TranslationQueue interface {
	AddTask(task Task)
	Start()
}

type Task struct {
	TranslationID TranslationID
	Text          string
}
