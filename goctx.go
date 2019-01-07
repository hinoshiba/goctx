package goctx

import (
	"context"
	"sync"
)

type Owner struct {
	Ctx	context.Context
	Cancel	context.CancelFunc
	Mux	sync.Mutex
}

type Worker struct {
	Ctx	context.Context
	Cancel	context.CancelFunc
	Mux	*sync.Mutex
}

func (self *Owner)NewWorker() Worker {
	ctx, _ := context.WithCancel(self.Ctx)
	return Worker{Ctx:ctx, Cancel:self.Cancel, Mux:&self.Mux}
}

func NewOwner() Owner {
	ctx, cancel := context.WithCancel(context.Background())
	return Owner{Ctx:ctx, Cancel:cancel, Mux:sync.Mutex{}}
}
