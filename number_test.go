package test

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(i int) {
	fmt.Println("number:", i)
}

func TestRunDisplayNumber(t *testing.T) {
	for i := range 100_000 {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
