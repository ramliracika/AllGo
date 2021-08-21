package goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string) // mambuat channel
	defer close(channel)         // harus close setelah selesai running channel

	go func() {
		channel <- "Ramli"  // input data ke channel
		channel <- "Racika" // input kedua channel
		fmt.Println("Selesai Input data ke channel")
		time.Sleep(2 * time.Millisecond)

	}()

	data := <-channel //mengambil data dari channel
	fmt.Println(data)

	result := <-channel // jumlah inputan === mengambil data
	fmt.Println(result)
	time.Sleep(1 * time.Millisecond)
}

func TestChannelParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go ChannelParameter(channel) //channel dengan parameter

	data := <-channel //mengambil data
	fmt.Println(data)

	time.Sleep(1 * time.Second)
}

func TestInOut(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)  //channel sebagai pengirim
	go OnlyOut(channel) //channel sebagai penerima

	time.Sleep(2 * time.Second)
	close(channel)

}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 4) //membuat kapasitas buffered
	defer close(channel)

	go BufferedIn(channel)
	go BufferedOut(channel)

	time.Sleep(2 * time.Second)
	fmt.Println("All Complete")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {

		for i := 0; i <= 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i) // konversi int ke string

		}
		close(channel)
	}()

	for data := range channel { //membuat perulangan range untuk channel
		fmt.Println("Menerima Data ", data)
	}

	fmt.Println("All Complete")
}

func TestSelect(t *testing.T) {
	counter := 0
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go ChannelParameter(channel1)
	go ChannelParameter2(channel2)

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++

		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++

		default:
			fmt.Println("Loading...")

		}
		if counter == 4 {
			break
		}
	}
	time.Sleep(1 * time.Second)
}

func TestRaceCondition(t *testing.T) {

	x := 0

	for i := 0; i < 100; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Total Counter : ", x)
}
