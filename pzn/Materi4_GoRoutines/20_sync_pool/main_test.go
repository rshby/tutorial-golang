package main

import (
	"fmt"
	"sync"
	"testing"
)

/*
Pool -> implementasi design pattern bernama object pool pattern
sederhananya, design pattern ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya kita bisa mengambil dari pool, dan selesai menggunakan kita bisa menyimpan kembali ke Poolnya
implementasi Pool di golang ini sudah aman dari problem race condition
*/

func TestPool(t *testing.T) {
	group := sync.WaitGroup{}
	pool := sync.Pool{}

	pool.Put("Reo")
	pool.Put("Sahobby")
	pool.Put("Reooo")

	for i := 0; i < 10; i++ {
		group.Wait()
		go func() {
			defer group.Done()

			group.Add(1)
			data := pool.Get() // -> mengambil data dari pool
			fmt.Println(data)
			pool.Put(data) // -> ditaruh ke pool lagi, karena setelah diGet() data akan hilang dari pool
		}()
	}

	group.Wait()
	fmt.Println("test complete!!!")
}
