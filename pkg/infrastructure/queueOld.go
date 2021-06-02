package infrastructure

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"text-to-speech-translation-service/pkg/domain"
)

// queueOld holds name, list of tasks and context with cancel.
type queueOld struct {
	name   string
	tasks  chan job
	ctx    context.Context
	cancel context.CancelFunc
}

// Worker responsible for queueOld serving.
type Worker struct {
	Queue *queueOld
}

type job struct {
	jobType int
	id      uuid.UUID
	data    map[string]string
}

// NewQueueOld instantiates new queueOld.
func NewQueueOld(name string) *queueOld {
	ctx, cancel := context.WithCancel(context.Background())

	return &queueOld{
		tasks:  make(chan job),
		name:   name,
		ctx:    ctx,
		cancel: cancel,
	}
}

// NewWorker initialises new Worker.
func NewWorker(queue *queueOld) *Worker {
	return &Worker{
		Queue: queue,
	}
}

// AddTask sends job to the channel.
func (q *queueOld) AddJob(task domain.Task) {
	go func(task domain.Task) {
		q.tasks <- job{id: uuid.UUID(task.ID), jobType: task.TaskType, data: task.Data}
	}(task)
	fmt.Printf("New job %s added to %s queueOld\n", task.ID, q.name)
}

// Run performs job execution.
func (j job) Run() error {
	fmt.Printf("Job %s is processed", j.data["name"])
	return nil
}

// DoWork processes tasks from the queueOld (tasks channel).
func (w *Worker) DoWork() bool {
	for {
		select {
		// if context was canceled.
		case <-w.Queue.ctx.Done():
			fmt.Printf("Work done in queueOld %s: %s!", w.Queue.name, w.Queue.ctx.Err())
			return true
		// if job received.
		case task := <-w.Queue.tasks:
			err := task.Run()
			if err != nil {
				fmt.Print(err)
				continue
			}
		}
	}
}
