package goroutine

import (
	"fmt"
	"sync"
	"time"
)

type UserBank struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBank) Lock() {
	user.Mutex.Lock()
}

func (user *UserBank) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBank) Change(amount int) {
	user.Balance = user.Balance + amount
}

func (user *UserBank) Transfer(user1 *UserBank, user2 *UserBank, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(2 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)
	time.Sleep(2 * time.Second)

	user1.Unlock()
	user2.Unlock()
}
