package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				mutex.Lock() //mengunci eksekusi
				x = x + 1
				mutex.Unlock() // melepas kunci
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Total Counter : ", x)

}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				account.AddBalance(1)
				fmt.Println("Total Penambahan : ", account.GetBalance())
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Total Balance : ", account.GetBalance())
}
