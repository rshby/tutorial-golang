package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Timer -> representasi suatu kejadian
Ketika waktu timer sudah expire, maka event akan dikirim ke dalam channel
untuk membuat Timer kita bisa menggunakan time.NewTimer(duration)
*/

func TestTimer(t *testing.T) {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

/*
time.After()
- kadang kita hanya butuh channelnya saja, tidak membutuhkan data Timernya
- untuk melakukan  hal itu bisa menggunakan function time.After(duration)
*/

func TestTimerAfter(t *testing.T) {
	channel := time.After(2 * time.Second)

	tick := <-channel
	fmt.Println(tick)
}

/*
time.AfterFunc()
- kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
- kita memanfaatkan Timer dengan menggunakan function time.AfterFunc()
- kita tidak perlu lagi mengguankan channelnya, cukup kirimkan function yang akan dipanggil ketika timer mengirim kejadian
*/

func TestAfterFunc(t *testing.T) {
	group := &sync.WaitGroup{}
	group.Add(1)

	fmt.Println(time.Now())
	time.AfterFunc(2*time.Second, func() {
		fmt.Println(time.Now())
		fmt.Println("Execute after 2 second")
		group.Done()
	})

	group.Wait()
	fmt.Println("Test Complete!!!")
}
