package tests

import (
	"testing"

	"unsafe"

	"github.com/xiaonanln/go-xnsyncutil/xnsyncutil"
)

func TestAtomicPointer(t *testing.T) {
	var ai xnsyncutil.AtomicPointer
	var a int
	ai.Store((unsafe.Pointer(&a)))
	*(*int)(ai.Load()) = 1
	if a != 1 {
		t.Errorf("a should be 1, but is %d", a)
	}
	b := 100

	ai.Store((unsafe.Pointer(&b)))
	*(*int)(ai.Load()) = 200
	if b != 200 {
		t.Errorf("b should be 200, but is %d", b)
	}

	if unsafe.Pointer(&b) != ai.Load() {
		t.Errorf("&b is %p, but load %p", &b, ai.Load())
	}
	//ai.Store(&b)
}
