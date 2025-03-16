package test

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(1 * time.Second)

		// goroutine akan blocking, jika tidak ada yang mengambil data dari channel
		channel <- "zukron"

		fmt.Println("data channel sudah diambil")
	}()

	name := <-channel
	fmt.Println(name)

	time.Sleep(5 * time.Second)
}
