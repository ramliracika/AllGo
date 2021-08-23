package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	data.Store(value, value) //menambahkan data kedalam map
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		go AddMap(data, i, group)
	}

	group.Wait()
	data.Range(func(key, value interface{}) bool { // untuk melakukan keseluruhan data di map
		fmt.Println(key, ":", value)
		return true
	})

}

// loadkey(key) untuk mengambil data dari map menggunakan key
// delete(key) untuk menghapus data
