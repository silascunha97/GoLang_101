package main

import (
	"fmt"
	"strconv"
)

//func hello(nome string) {
//	fmt.Println("Olá", nome, "❤")
//}

//func soma(val1, val2 int) int {
//	return val1 + val2
//}

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	total = a + i
	return
}

func main() {
	//hello("Silas Augusto a soma é")
	//fmt.Println(soma(10, 5))
	total, err := convertAndSum(10, "23")
	fmt.Println("total:", total, err)
}
