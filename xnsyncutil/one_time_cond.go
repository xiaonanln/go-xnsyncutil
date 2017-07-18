package xnsyncutil

import "sync"

type OneTimeCond struct {
	signalled bool
	lock      sync.Mutex
	subcond   *sync.Cond
}

func NewOneTimeCond() *OneTimeCond {
	cond := &OneTimeCond{}
	cond.subcond = sync.NewCond(&cond.lock)
	return cond
}

func (cond *OneTimeCond) Wait() {
	cond.lock.Lock()
	if cond.signalled {
		cond.lock.Unlock()
		return
	}

	cond.subcond.Wait()
	cond.lock.Unlock()
}

func (cond *OneTimeCond) Signal() {
	cond.lock.Lock()
	cond.signalled = true
	cond.subcond.Broadcast()
	cond.lock.Unlock()
}

func (cond *OneTimeCond) IsSignalled() (signalled bool) {
	cond.lock.Lock()
	signalled = cond.signalled
	cond.lock.Unlock()
	return
}
