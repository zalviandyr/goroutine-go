package test

import (
	"fmt"
	"sync"
	"testing"
)

func DoPrint(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	fmt.Println("Angka:", i)
}

func TestCreateWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go DoPrint(&wg, i)
	}

	wg.Wait()
}
