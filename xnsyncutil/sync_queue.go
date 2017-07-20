package xnsyncutil

import (
	"sync"

	"gopkg.in/eapache/queue.v1"
)

// Synchronous FIFO queue
type SyncQueue struct {
	lock    sync.Mutex
	popable *sync.Cond
	buffer  *queue.Queue
}

// Create a new SyncQueue
func NewSyncQueue() *SyncQueue {
	ch := &SyncQueue{
		buffer: queue.New(),
	}
	ch.popable = sync.NewCond(&ch.lock)
	return ch
}

// Pop an item from SyncQueue, will block if SyncQueue is empty
func (q *SyncQueue) Pop() interface{} {
	c := q.popable
	buffer := q.buffer

	q.lock.Lock()
	for buffer.Length() == 0 {
		c.Wait()
	}

	v := buffer.Peek()
	buffer.Remove()

	q.lock.Unlock()
	return v
}

// Try to pop an item from SyncQueue, will return immediately with bool=false if SyncQueue is empty
func (q *SyncQueue) TryPop() (interface{}, bool) {
	buffer := q.buffer

	q.lock.Lock()

	if buffer.Length() > 0 {
		v := buffer.Peek()
		buffer.Remove()
		q.lock.Unlock()
		return v, true
	} else {
		q.lock.Unlock()
		return nil, false
	}
}

// Push an item to SyncQueue. Always returns immediately without blocking
func (q *SyncQueue) Push(v interface{}) {
	q.lock.Lock()
	q.buffer.Add(v)
	q.popable.Signal()
	q.lock.Unlock()
}

// Get the length of SyncQueue
func (q *SyncQueue) Len() (l int) {
	q.lock.Lock()
	l = q.buffer.Length()
	q.lock.Unlock()
	return
}
