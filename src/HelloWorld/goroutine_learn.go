package HelloWorld

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0

	for {
		i++
		fmt.Printf("newTask: i = %d \n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("mainTask: i = %d \n", i)
		time.Sleep(time.Second)
	}
}
