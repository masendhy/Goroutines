package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// create channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- " sendhe boedhi"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "masendhy"
}

//channel as a parameter
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// adding buffer at channel
func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "sendhy"
		channel <- "boedhi"
		channel <- "satriya"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Done")

}

// Range channel : untuk menerima data dengan perulangan

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- "Perulangan ke : " + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("OK")
}

// select channel

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1 :", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2 :", data)
			counter++
			// menambahkan default untuk mengetahui progrma kita masih berjaln dan menunggu data masuk

		default:
			fmt.Println("Data Belum Masuk")

		}
		if counter == 2 {
			break
		}
	}

}
