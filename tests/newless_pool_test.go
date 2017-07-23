package tests

import (
	"testing"

	"github.com/xiaonanln/go-xnsyncutil/xnsyncutil"
)

func TestNewNewlessPool(t *testing.T) {
	pool := xnsyncutil.NewNewlessPool()
	if pool == nil {
		t.Errorf("pool is nil")
	}
}

func TestNewlessPool_Basic(t *testing.T) {
	pool := xnsyncutil.NewNewlessPool()
	pool.Put(1)
	if pool.Get().(int) != 1 {
		t.Errorf("put 1 but get not 1")
	}

	if pool.TryGet() != nil {
		t.Errorf("should be nil")
	}

	pool.Put(3)
	if pool.TryGet().(int) != 3 {
		t.Error("wrong")
	}

	go func() {
		if pool.Get().(int) != 2 {
			t.Errorf("wrong")
		}
	}()
	pool.Put(2)
}
