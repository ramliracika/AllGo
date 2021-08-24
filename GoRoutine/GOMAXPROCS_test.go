package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomax(t *testing.T) {
	cpu := runtime.NumCPU()
	fmt.Println("Total CPU :", cpu)

	// runtime.GOMAXPROCS(20) --> untuk menambah jumlah thread
	thread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", thread)

	goroutine := runtime.NumGoroutine()
	fmt.Println("Total GoRoutine :", goroutine)
}
