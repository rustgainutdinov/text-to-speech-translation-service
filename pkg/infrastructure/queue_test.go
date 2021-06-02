package infrastructure

import (
	"fmt"
	"testing"
)

func TestQueue_AddTask(t *testing.T) {
	queue := NewQueue()
	go queue.start()
	queue.AddTask("msg 1")
	queue.AddTask("msg 2")
	queue.AddTask("msg 3")
	queue.AddTask("msg 4")
	fmt.Println("end of test")
}
