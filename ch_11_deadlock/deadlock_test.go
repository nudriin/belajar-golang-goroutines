package ch_11_deadlock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBallance struct {
	Mutex    sync.Mutex
	Name     string
	Ballance int
}

func (user *UserBallance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBallance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBallance) Change(amount int) {
	user.Ballance = user.Ballance + amount
}

func Transfer(user1 *UserBallance, user2 *UserBallance, amount int) {
	user1.Lock()
	fmt.Println("Lock usr1:", user1.Name)
	user1.Change(-amount) // mengurangi ballance

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock usr2:", user2.Name)
	user2.Change(amount) // menambah ballance

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBallance{
		Name:     "Elon Musk",
		Ballance: 10,
	}

	user2 := UserBallance{
		Name:     "Mark",
		Ballance: 10,
	}

	go Transfer(&user1, &user2, 2)
	go Transfer(&user2, &user1, 3) // * user 1 tidak bertambah

	time.Sleep(5 * time.Second)

	fmt.Printf("%s ballance %d\n", user1.Name, user1.Ballance)
	fmt.Printf("%s ballance %d\n", user2.Name, user2.Ballance)

}
