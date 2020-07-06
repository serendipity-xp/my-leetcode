package other

import (
	"fmt"
	"sync"
)

func Locker()  {
	defer func() {
		fmt.Println("Try to recover the panic")
		if p := recover(); p != nil{
			fmt.Println("recover the panic:", p)
		}
	}()
	var mutex sync.Mutex
	fmt.Println("begin lock")
	mutex.Lock()
	fmt.Println("get locked")
	fmt.Println("unlock lock")
	mutex.Unlock()
	fmt.Println("lock is unlocked")
	fmt.Println("unlock lock again")
	mutex.Unlock()
}