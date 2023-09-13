package main

import (
	"fmt"
	"sync"
	"testing"
)

/*
Once -> fitur di golang untuk memastikan sebuah function hanya dieksekusi sekali saja
jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi functionnya
goroutine yang lain akan dihiraukan, artinya function tidak akan diakses lagi
*/

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter) // -> hasilnya 1, karena functionnya sama jadi hanya dieksekusi sekali
}
