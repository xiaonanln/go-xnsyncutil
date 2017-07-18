package xnsyncutil

import "sync/atomic"

type AtomicInt int64

func (ai *AtomicInt) Store(v int) {
	atomic.StoreInt64((*int64)(ai), int64(v))
}

func (ai *AtomicInt) Add(delta int) {
	atomic.AddInt64((*int64)(ai), int64(delta))
}

func (ai *AtomicInt) Load() int {
	return int(atomic.LoadInt64((*int64)(ai)))
}

type AtomicBool AtomicInt

func (ab *AtomicBool) Store(v bool) {
	if v {
		(*AtomicInt)(ab).Store(1)
	} else {
		(*AtomicInt)(ab).Store(0)
	}
}

func (ab *AtomicBool) Load() bool {
	return (*AtomicInt)(ab).Load() != 0
}
