package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounterrr(ctx context.Context) chan int {
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
				time.Sleep(1 * time.Second) // slow simulation
			}
		}

	}()

	return channel
}

func TestWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	channel := CreateCounterr(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for i := range channel {
		fmt.Println("Counter", i)
		if i == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
