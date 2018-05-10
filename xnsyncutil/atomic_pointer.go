package xnsyncutil

import (
	"unsafe"
	"sync/atomic"
)

type AtomicPointer struct {
	ptr unsafe.Pointer
}

// Store ptr value atomically
func (ap *AtomicPointer) Store(pointer unsafe.Pointer) {
	atomic.StorePointer(&ap.ptr, pointer)
}


// Load ptr value atomically
func (ap *AtomicPointer) Load() unsafe.Pointer {
	return atomic.LoadPointer(&ap.ptr)
}

