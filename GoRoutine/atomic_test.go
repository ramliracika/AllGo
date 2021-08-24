package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// atomic berfungsi untuk handle data primitive, dab sudah bebas dari race conditio
// dan berfungsi seperti mutex untuk locking

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {

		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1) // x=x+1
			}
			group.Done()

		}()

	}
	group.Wait()
	fmt.Println("Counter : ", x)

}
