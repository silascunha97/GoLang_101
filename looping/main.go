package main

import "fmt"

func tabuada_1() {
	/*Fácil  1 – Faça um programa que receba um número
	e usando laços de repetição calcule e mostre a tabuada
	desse número. */

	var numero int

	fmt.Println("Digite um número: ")
	fmt.Scanln(&numero)

	fmt.Println("Tabuada do ", numero)
	for i := 0; i <= 10; i++ {
		fmt.Println(numero, "X", i, " = ", numero*i)
	}
}

func tabuada_2() {
	/*
		Fácil  2 –  Faça um programa que mostre as tabuadas dos
		 números de 1 a 10 usando laços de repetição.
	*/
	for i := 1; i <= 10; i++ {
		fmt.Println("Tabuada do ", i)
		for j := 0; j <= 10; j++ {
			fmt.Println(i, " x ", j, " = ", i*j)
		}
		fmt.Println("##########################")
	}
}

func main() {
	//nomes := []string{"Arthur", "Amanda", "Felipe", "João", "Joaquim do verde"}
	//for i := 0; i < len(nomes); i++ {
	//	fmt.Println(nomes[i])
	//}

	//for i, nome := range nomes {
	//	fmt.Println(i, nome)
	//}
	//var i int
	//for i, nome := range nomes {
	//	fmt.Println(i, nome)
	//}

	//nomes = append(nomes, "Augusto")
	//for i, nome := range nomes {
	//	fmt.Println(i, nome)
	//}

	/*for i := 0; i < len(nomes); i++ {
		fmt.Println(nomes[i])
	} */

	//fmt.Println(nomes, len(nomes), cap(nomes))

	//nomes := make([]string, 0)
	//nomes = append(nomes, "Silas")
	tabuada_2()
}
