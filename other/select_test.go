package other

import (
	"fmt"
	"testing"
)

/**
* @DateTime   : 2020/7/23 17:22
* @Author     : xulp
* @Description:
**/

func TestSelect(t *testing.T) {
	t.Run("", func(t *testing.T) {
		ch := make(chan int, 2)
		for i := 0; i < 10; i++ {
			select {
			case x := <-ch:
				fmt.Println(x) // "0" "2" "4" "6" "8"
			case ch <- i:
			}
		}
	})
}