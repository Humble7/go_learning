package share_memory

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	count := 0
	for i := 0; i < 5000; i ++ {
		go func() {
			count ++
		}()
	}

	time.Sleep(time.Second * 1)
	t.Log(count)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	count := 0
	for i := 0; i < 5000; i ++ {
		go func() {
			defer func() {
                mut.Unlock()
			}()
			mut.Lock()
			count ++
		}()
	}

	time.Sleep(time.Second * 1)
	t.Log(count)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5000; i ++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count ++
			wg.Done()
		}()
	}

	wg.Wait()
	t.Log(count)
}
