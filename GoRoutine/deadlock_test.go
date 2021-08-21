package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestDeadLock(t *testing.T) {
	user1 := UserBank{
		Name:    "Ramli",
		Balance: 10,
	}

	user2 := UserBank{
		Name:    "Haryadi",
		Balance: 10,
	}
	user := UserBank{}

	go user.Transfer(&user1, &user2, 5)
	// go user.Transfer(&user2, &user1, 10) // will deadlock

	time.Sleep(3 * time.Second)

	fmt.Println("User 1 : ", user1.Name, " Balance : ", user1.Balance)

	fmt.Println("User 2 : ", user2.Name, " Balance : ", user2.Balance)
}
