package infrastructure

import (
	"context"
	"fmt"
	"time"
)

type queue struct {
	in   chan interface{}
	stop chan struct{}
	ctx  context.Context
}

func (s queue) AddTask(msg interface{}) {
	go func() {
		s.in <- msg
	}()
}

func (s queue) Close() {
	close(s.stop)
}

func (s queue) start() {
	for {
		select {
		case msg := <-s.in:
			{
				fmt.Println(msg)
			}
		case <-s.stop:
			{
				close(s.in)
				return
			}
		default:
			fmt.Println("default")
			time.Sleep(2 * time.Second)
		}
	}
}

func NewQueue() queue {
	sub := queue{
		in:   make(chan interface{}),
		stop: make(chan struct{}),
		ctx:  context.Background(),
	}
	return sub
}
