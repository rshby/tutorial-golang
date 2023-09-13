package deadlock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
kejadian misal kita salah implementasi mutex atau locking
- deadlock adalah keadaan dimana sebuah proses goroutines saling menunggu lock sehingga tidak ada satupun goroutine yang berjalan
*/

// buat struct userBalance
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

// function untuk melakukan lock pada object UserBalance
func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

// buat function untuk melakukan UnLock pada object UserBalance
func (user *UserBalance) UnLock() {
	user.Mutex.Unlock()
}

// buat function untuk merubah balance
func (user *UserBalance) Change(inputAmount int) {
	user.Balance += inputAmount
}

// buat function untuk melakukan Transfer
func Transfer(inputUser1 *UserBalance, inputUser2 *UserBalance, inputAmount int) {
	inputUser1.Lock()
	fmt.Println("Lock", inputUser1.Name)
	inputUser1.Change(-inputAmount)

	time.Sleep(5 * time.Millisecond)

	inputUser2.Lock()
	fmt.Println("Lock", inputUser2.Name)
	inputUser2.Change(inputAmount)

	time.Sleep(5 * time.Millisecond)

	inputUser1.Unlock()
	inputUser2.UnLock()
}

// buat function untuk mengetest Deadlock
func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Reo",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Eko",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 1)
	time.Sleep(2 * time.Second)
	go Transfer(&user2, &user1, 1)
	time.Sleep(2 * time.Second)

	fmt.Println("User", user1.Name, "Balance", user1.Balance)
	fmt.Println("User", user2.Name, "Balance", user2.Balance)

}
