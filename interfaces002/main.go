package main

import (
	"fmt"
	"time"
)

func numeros() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n ", i)
		time.Sleep(1 * time.Millisecond * 150)
	}
}

func letras() {
	for i := 'a'; i < 'a'+26; i++ {
		fmt.Printf("%c \n", i)
		time.Sleep(1 * time.Millisecond * 150)
	}
}

func main() {
	//numeros()
	letras()
}
