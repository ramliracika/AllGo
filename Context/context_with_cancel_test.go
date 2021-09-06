package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func IncrementCounter() chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		counter := 1
		for {
			channel <- counter
			counter++
		}

	}()

	return channel
}

func TestWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	channel := IncrementCounter()

	for i := range channel {
		fmt.Println("Counter", i)
		if i == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine()) // goroutine leak, still exist even program done

}

//--------------------------

func CreateCounter(ctx context.Context) chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				channel <- counter
				counter++

			}
		}

	}()

	return channel
}

func TestWithCan2(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	channel := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for i := range channel {
		fmt.Println("Counter", i)
		if i == 10 {
			break
		}
	}

	cancel() // send signal cancel to context
	time.Sleep(5 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
