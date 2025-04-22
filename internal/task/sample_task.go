package task

import (
	"context"
	"fmt"
	"time"
)

type SampleTask struct {
	id    string
	input string
	delay time.Duration
}

func NewSampleTask(input string, delay time.Duration) *SampleTask {
	return &SampleTask{
		id:    fmt.Sprintf("sample-%d", time.Now().UnixNano()),
		input: input,
		delay: delay,
	}
}

func (t *SampleTask) ID() string {
	return t.id
}

func (t *SampleTask) Execute(ctx context.Context) (string, error) {
	time.Sleep(t.delay) // имитируем долгую I/O задачу
	return fmt.Sprintf("Processed input: %s", t.input), nil
}
