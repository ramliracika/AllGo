package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{ // membuat pool
		New: func() interface{} { // menambah default value
			return "New"
		},
	}

	pool.Put("Ramli") // put untuk menambahkan data
	pool.Put("Racika")
	pool.Put("Haryadi")
	pool.Put("Maulana")

	for i := 0; i < 10; i++ {
		go func() {

			data := pool.Get() // mengambil data
			fmt.Println(data)  // menampilkan isi pool

			time.Sleep(1 * time.Second)
			pool.Put(data) // menambahkan data kedalam pool

		}()
	}

	time.Sleep(2 * time.Second)
}
