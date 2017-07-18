package tests

import (
	"testing"

	"sync"

	"time"

	"github.com/xiaonanln/go-xnsyncutil/xnsyncutil"
)

func TestAtomicInt(t *testing.T) {
	var ai xnsyncutil.AtomicInt
	if ai.Load() != 0 {
		t.Errorf("should be initialized to 0")
	}
	ai.Store(100)
	if ai.Load() != 100 {
		t.Errorf("load wrong value")
	}
	ai.Add(100)
	if ai.Load() != 200 {
		t.Errorf("load wrong value")
	}

	var wait sync.WaitGroup
	concurrent := 2
	wait.Add(concurrent)

	for i := 0; i < concurrent; i++ {
		var delta int
		if i%2 == 0 {
			delta = -1
		} else {
			delta = 1
		}
		go func() {
			for i := 0; i < 1000; i++ {
				ai.Add(delta)
				time.Sleep(time.Microsecond)
			}
			wait.Done()
		}()
	}

	wait.Wait()
	if ai.Load() != 200 {
		t.Errorf("load wrong value after fuzzy concurrent add")
	}
}

func TestAtomicBool(t *testing.T) {
	var ab xnsyncutil.AtomicBool
	if ab.Load() {
		t.Errorf("should be initialized to false")
	}
	ab.Store(true)
	if ab.Load() != true {
		t.Errorf("load wrong value")
	}
	ab.Store(false)
	if ab.Load() != false {
		t.Errorf("load wrong value")
	}

	var wait sync.WaitGroup
	concurrent := 2
	wait.Add(concurrent)

	for i := 0; i < concurrent; i++ {
		val := i%2 == 0
		go func() {
			for i := 0; i < 1000; i++ {
				ab.Store(val)
				time.Sleep(time.Microsecond)
			}
			wait.Done()
		}()
	}

	wait.Wait()
}
