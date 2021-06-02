package infrastructure

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
	"text-to-speech-translation-service/pkg/domain"
)

var products = []string{
	"books",
	"computers",
}

func TestWorker_DoWork(t *testing.T) {
	newProducts := []string{
		"apples",
		"oranges",
		"wine",
		"bread",
		"orange juice",
	}

	productsQueue := NewQueueOld("NewProducts")
	for _, newProduct := range newProducts {
		productsQueue.AddJob(domain.Task{
			TaskType: domain.TaskTypeSpeechToTextTranslation,
			ID:       uuid.New(),
			Data:     map[string]string{"name": newProduct},
		})
	}

	worker := NewWorker(productsQueue)
	worker.DoWork()
	defer fmt.Print(products)
}

func TestQueue_AddJob(t *testing.T) {
	fmt.Printf("qq")
}
