package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}        // sebagai kunci cond
var cond = sync.NewCond(&locker) // deklarasi cond dengan kunci
var group = sync.WaitGroup{}

func InputData(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()              //locking
	cond.Wait()                // menunggu signal/broadcast untuk next executions
	fmt.Println("Done", value) // menampilkan data
	cond.L.Unlock()            // open lock
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ { // iterasi func
		go InputData(i)
	}

	go func() { // iterasi untuk mengirim signal
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // signal untuk memberi izin wait untuk eksekusi satu persatu
		}

	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast() // broadcast mengizinkan semua wait berjalan
	// }()

	group.Wait()
}
