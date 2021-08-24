package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second) // deklarasi waktu timer
	fmt.Println(time.Now())

	data := <-timer.C
	fmt.Println(data) // akan dieksekusi saat timer expired
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second) // deklarasi after
	fmt.Println(time.Now())

	data := <-channel // langsung memanggil channel
	fmt.Println(data)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() { // menjalankan function setelah timer expired
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now()) // tidak membutuhkan channel
	group.Wait()
}
