package goroutine

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go SayHai(group)
	}
	group.Wait()

}
