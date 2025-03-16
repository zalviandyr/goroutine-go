package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

type Counter struct {
	sync.Mutex
	val int
}

func (c *Counter) Add(int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *Counter) Value() int {
	return c.val
}

func TestMutex(t *testing.T) {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter Counter

	for range 10_000 {
		wg.Add(1)

		go func() {
			for range 10_000 {
				meter.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("meter value:", meter.Value())
}
