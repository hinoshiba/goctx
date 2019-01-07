package goctx

import (
	"context"
	"sync"
)

type Owner struct {
	ctx	context.Context
	cancel	context.CancelFunc
	mux	sync.Mutex
}

type Worker struct {
	ctx	context.Context
	cancel	context.CancelFunc
	mux	*sync.Mutex
}

func NewOwner() Owner {
	ctx, cancel := context.WithCancel(context.Background())
	return Owner{ctx:ctx, cancel:cancel, mux:sync.Mutex{}}
}

func (self *Owner)NewWorker() Worker {
	ctx, _ := context.WithCancel(self.ctx)
	return Worker{ctx:ctx, cancel:self.cancel, mux:&self.mux}
}

func (self *Owner) Lock() {
	self.mux.Lock()
}

func (self *Owner) Unlock() {
	self.mux.Unlock()
}

func (self *Owner) Cancel() {
	self.cancel()
}

func (self *Worker) Lock() {
	self.mux.Lock()
}

func (self *Worker) Unlock() {
	self.mux.Unlock()
}

func (self *Worker) Cancel() {
	self.cancel()
}

func (self *Worker) RecvCancel() <-chan struct{} {
	return self.ctx.Done()
}
