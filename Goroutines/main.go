package main

import (
	"fmt"
	"time"
)

func numero() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func letras() {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c \n", l)
		time.Sleep(time.Millisecond * 250)
	}
}

func main() {
	go numero()
	go letras()
}
