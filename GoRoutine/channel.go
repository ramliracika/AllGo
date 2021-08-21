package goroutine

import (
	"fmt"
	"time"
)

func ChannelParameter(channel chan string) {
	time.Sleep(1 * time.Millisecond)

	channel <- "Ramli"
	channel <- "Racika"
}
func ChannelParameter2(channel chan string) {
	time.Sleep(1 * time.Millisecond)

	channel <- "Haryadi"
	channel <- "Maulana"
}

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Raka Hana"
}

func OnlyOut(channel <-chan string) {

	data := <-channel
	fmt.Println(data)
}

func BufferedIn(channel chan string) {
	time.Sleep(1 * time.Second)

	channel <- "Ramli"
	channel <- "Racika"
	channel <- "Haryadi"
	channel <- "Maulana"

}

func BufferedOut(channel chan string) {
	time.Sleep(1 * time.Second)

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
