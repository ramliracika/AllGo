package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGoHello(t *testing.T) {
	go HelloWorld()

	fmt.Println("after goroutine")

	time.Sleep(1 * time.Second)
}

func TestDisplay(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
