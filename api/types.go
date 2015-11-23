package api

import (
	"fmt"

	"golang.org/x/net/context"
)

type Item interface {
	Key() string
	Data() interface{}
}

type Process interface {
	GetName() string
	Exec(context.Context) error
	Init(context.Context) error
	Uninit(context.Context) error
}

type Source interface {
	GetOutput() <-chan interface{}
}

type Sink interface {
	SetInput(<-chan interface{})
}

type Processor interface {
	Process
	Source
	Sink
}

type Endpoint interface {
	Done() <-chan struct{}
}

type Collector interface {
	SetInputs([]<-chan interface{})
}

type Emitter interface {
	GetOutputs() []<-chan interface{}
}

type ProcError struct {
	Err      error
	ProcName string
}

func (e ProcError) Error() string {
	if e.ProcName != "" {
		return fmt.Sprintf("[%s] %v", e.ProcName, e.Err)
	}
	return e.Err.Error()
}
