package tests

import (
	"testing"

	"time"

	"github.com/xiaonanln/go-xnsyncutil/xnsyncutil"
)

func TestOneTimeCond(t *testing.T) {
	cond := xnsyncutil.NewOneTimeCond()
	if cond.IsSignalled() {
		t.Errorf("should not be signalled when created")
	}

	startWaitTime := time.Now()

	go func() {
		time.Sleep(time.Millisecond * 100)
		cond.Signal()
	}()

	cond.Wait()
	if time.Now().Sub(startWaitTime) < time.Millisecond*90 {
		t.Errorf("wait return too soon")
	}

	cond.Wait() // should return immediately
}
