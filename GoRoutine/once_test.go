package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func NotifyMe() {

	counter++
}

func TestOnce(t *testing.T) {

	once := &sync.Once{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {

		go func() {
			group.Add(1)
			once.Do(NotifyMe)
			group.Done()
		}()

	}
	fmt.Println("counter", counter)
	group.Wait()

}
