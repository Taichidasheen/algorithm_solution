package practice

import (
	"fmt"
	"time"
)

func print100() {
	flag1 := make(chan int , 1)
	flag2 := make(chan int, 1)
	flag1 <- 1
	go print1(flag1, flag2)
	go print2(flag2, flag1)
	time.Sleep(2*time.Second)
}

func print1(flag1 chan int, flag2 chan int) {
	count := 1
	for _ = range flag1 {
		if count > 99 {
			break
		}
		fmt.Println(count)
		count = count + 2
		flag2 <- 1
	}
}

func print2(flag1 chan int, flag2 chan int) {
	count := 2
	for _ = range flag1 {
		if count > 100 {
			break
		}
		fmt.Println(count)
		count = count + 2
		flag2 <- 1

	}
}
