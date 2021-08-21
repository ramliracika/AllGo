package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func SayHai(group *sync.WaitGroup) {
	defer group.Done() // harus di tutup

	group.Add(1) // jumlah goroutine yang akan di execute

	fmt.Println("Hai Ramli")

	time.Sleep(1 * time.Second)

}
