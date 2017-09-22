package tests

import (
	"testing"

	"github.com/xiaonanln/go-xnsyncutil/xnsyncutil"

	"fmt"

	"sync"
)

func TestSpinLock(t *testing.T) {

	var lock xnsyncutil.SpinLock
	var val int
	const N = 10000
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		for i := 0; i < N; i++ {
			lock.Lock()
			val += 1
			lock.Unlock()
		}
		wait.Done()
	}()

	go func() {
		for i := 0; i < N; i++ {
			lock.Lock()
			val += 1
			lock.Unlock()
		}
		wait.Done()
	}()

	wait.Wait()
	if val != N*2 {
		t.Fatalf("val should be %d, but is %d", N*2, val)
	}
	fmt.Printf("%d\n", val)
}

func BenchmarkSpinLock(b *testing.B) {

	var lock xnsyncutil.SpinLock
	const N = 10000

	for i := 0; i < b.N; i++ {
		var val int
		var wait sync.WaitGroup
		wait.Add(2)

		go func() {
			for i := 0; i < N; i++ {
				lock.Lock()
				val += 1
				lock.Unlock()
			}
			wait.Done()
		}()

		go func() {
			for i := 0; i < N; i++ {
				lock.Lock()
				val += 1
				lock.Unlock()
			}
			wait.Done()
		}()

		wait.Wait()
	}
}
func BenchmarkMutex(b *testing.B) {

	var lock sync.Mutex
	const N = 10000

	for i := 0; i < b.N; i++ {
		var val int
		var wait sync.WaitGroup
		wait.Add(2)

		go func() {
			for i := 0; i < N; i++ {
				lock.Lock()
				val += 1
				lock.Unlock()
			}
			wait.Done()
		}()

		go func() {
			for i := 0; i < N; i++ {
				lock.Lock()
				val += 1
				lock.Unlock()
			}
			wait.Done()
		}()

		wait.Wait()
	}
}
